// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package ipxe provides iPXE functionality.
package ipxe

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/siderolabs/talos/pkg/machinery/constants"
	"go.uber.org/zap"
)

const (
	ipxeScriptTemplateFormat = `#!ipxe
chain --replace %s
`
	archArm64 = "arm64"
	archAmd64 = "amd64"

	// initScriptName is the name of the iPXE init script served by the HTTP server.
	//
	// Some UEFIs with built-in iPXE require the script URL to be in the form of a filename ending with ".ipxe", hence we serve it under this path.
	initScriptName = "init.ipxe"

	// bootScriptName is the name of the iPXE boot script served by the HTTP server.
	//
	// Some UEFIs with built-in iPXE require the script URL to be in the form of a filename ending with ".ipxe", hence we serve it under this path.
	bootScriptName = "boot.ipxe"
)

// ImageFactoryClient represents an image factory client which ensures a schematic exists on image factory, and returns the PXE URL to it.
type ImageFactoryClient interface {
	SchematicIPXEURL(ctx context.Context, talosVersion, arch string, extensions, extraKernelArgs []string) (string, error)
}

// HandlerOptions represents the options for the iPXE handler.
type HandlerOptions struct {
	APIAdvertiseAddress string
	TalosVersion        string
	ExtraKernelArgs     string
	Extensions          []string
	APIPort             int
}

// Handler represents an iPXE handler.
type Handler struct {
	imageFactoryClient ImageFactoryClient
	logger             *zap.Logger
	kernelArgs         []string
	initScript         []byte
	options            HandlerOptions
}

// ServeHTTP serves the iPXE request.
//
// URL pattern: http://ip-of-this-server:50042/ipxe/boot.ipxe?uuid=${uuid}&mac=${net${idx}/mac:hexhyp}&domain=${domain}&hostname=${hostname}&serial=${serial}&arch=${buildarch}
//
// Implements http.Handler interface.
func (handler *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.PathValue("script") {
	default:
		handler.logger.Error("invalid iPXE script", zap.String("script", req.PathValue("script")))

		w.WriteHeader(http.StatusNotFound)

		return
	case initScriptName:
		handler.handleInitScript(w)

		return
	case bootScriptName:
	}

	ctx := req.Context()
	query := req.URL.Query()
	uuid := query.Get("uuid")
	mac := query.Get("mac")
	arch := query.Get("arch")
	logger := handler.logger.With(zap.String("uuid", uuid), zap.String("mac", mac), zap.String("arch", arch))

	if arch != archArm64 { // https://ipxe.org/cfg/buildarch
		arch = archAmd64 // qemu comes as i386, but we still want to boot amd64
	}

	logger.Info("handle iPXE boot request")

	// TODO: later, we can do per-machine kernel args and system extensions here

	kernelArgs := slices.Concat(handler.kernelArgs, handler.consoleKernelArgs(arch))

	body, statusCode, err := handler.bootViaFactoryIPXEScript(ctx, handler.options.TalosVersion, arch, handler.options.Extensions, kernelArgs)
	if err != nil {
		handler.logger.Error("failed to get iPXE script", zap.Error(err))

		w.WriteHeader(http.StatusInternalServerError)

		if _, err = w.Write([]byte("failed to get iPXE script: " + err.Error())); err != nil {
			handler.logger.Error("failed to write error response", zap.Error(err))
		}

		return
	}

	w.WriteHeader(statusCode)

	if _, err = w.Write([]byte(body)); err != nil {
		handler.logger.Error("failed to write response", zap.Error(err))
	}
}

func (handler *Handler) handleInitScript(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")

	if _, err := w.Write(handler.initScript); err != nil {
		handler.logger.Error("failed to write init script", zap.Error(err))
	}
}

func (handler *Handler) bootViaFactoryIPXEScript(ctx context.Context, talosVersion, arch string, extensions, kernelArgs []string) (body string, statusCode int, err error) {
	ipxeURL, err := handler.imageFactoryClient.SchematicIPXEURL(ctx, talosVersion, arch, extensions, kernelArgs)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("failed to get schematic IPXE URL: %w", err)
	}

	ipxeScript := fmt.Sprintf(ipxeScriptTemplateFormat, ipxeURL)

	return ipxeScript, http.StatusOK, nil
}

func (handler *Handler) consoleKernelArgs(arch string) []string {
	switch arch {
	case archArm64:
		return []string{"console=tty0", "console=ttyAMA0"}
	default:
		return []string{"console=tty0", "console=ttyS0"}
	}
}

// NewHandler creates a new iPXE server.
func NewHandler(configServerEnabled bool, imageFactoryClient ImageFactoryClient, options HandlerOptions, logger *zap.Logger) (*Handler, error) {
	initScript, err := buildInitScript(options.APIAdvertiseAddress, options.APIPort)
	if err != nil {
		return nil, fmt.Errorf("failed to build init script: %w", err)
	}

	logger.Info("patch iPXE binaries")

	if err = patchBinaries(initScript, logger); err != nil {
		return nil, err
	}

	logger.Info("successfully patched iPXE binaries")

	kernelArgs := strings.Fields(options.ExtraKernelArgs)

	if configServerEnabled {
		apiHostPort := net.JoinHostPort(options.APIAdvertiseAddress, strconv.Itoa(options.APIPort))
		talosConfigURL := fmt.Sprintf("http://%s/config?u=${uuid}", apiHostPort)

		kernelArgs = append(kernelArgs, fmt.Sprintf("%s=%s", constants.KernelParamConfig, talosConfigURL))
	}

	return &Handler{
		imageFactoryClient: imageFactoryClient,
		options:            options,
		kernelArgs:         kernelArgs,
		initScript:         initScript,
		logger:             logger,
	}, nil
}
