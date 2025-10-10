// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
	EnsureSchematic(ctx context.Context, extensions, extraKernelArgs []string) (string, error)
	GetIPXEURL(schematicID, talosVersion, arch string) (string, error)
}

// HandlerOptions represents the options for the iPXE handler.
type HandlerOptions struct {
	APIAdvertiseAddress string
	TalosVersion        string
	ExtraKernelArgs     string
	SchematicID         string
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

	consoleKernelArgs := handler.consoleKernelArgs(arch)
	kernelArgs := slices.Concat(handler.kernelArgs, consoleKernelArgs)

	logger.Debug("injected console kernel args to the iPXE request", zap.Strings("console_kernel_args", consoleKernelArgs))

	body, statusCode, err := handler.bootViaFactoryIPXEScript(ctx, arch, kernelArgs)
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

func (handler *Handler) bootViaFactoryIPXEScript(ctx context.Context, arch string, kernelArgs []string) (body string, statusCode int, err error) {
	schematicID := handler.options.SchematicID

	if schematicID == "" {
		if schematicID, err = handler.imageFactoryClient.EnsureSchematic(ctx, handler.options.Extensions, kernelArgs); err != nil {
			return "", http.StatusInternalServerError, fmt.Errorf("failed to get schematic IPXE URL: %w", err)
		}
	}

	var ipxeURL string

	if ipxeURL, err = handler.imageFactoryClient.GetIPXEURL(schematicID, handler.options.TalosVersion, arch); err != nil {
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
func NewHandler(ctx context.Context, configServerEnabled bool, imageFactoryClient ImageFactoryClient, options HandlerOptions, logger *zap.Logger) (*Handler, error) {
	apiHostPort := net.JoinHostPort(options.APIAdvertiseAddress, strconv.Itoa(options.APIPort))
	talosConfigURL := fmt.Sprintf("http://%s/config?u=${uuid}", apiHostPort)
	talosConfigKernelArg := fmt.Sprintf("%s=%s", constants.KernelParamConfig, talosConfigURL)

	if options.SchematicID != "" {
		if len(options.Extensions) > 0 || len(options.ExtraKernelArgs) > 0 {
			return nil, fmt.Errorf("schematicID cannot be used with extensions or extraKernelArgs")
		}

		if configServerEnabled {
			logger.Sugar().Warnf("schematic ID is set explicitly to %q and the config server is enabled (e.g., Omni connection is requested), "+
				"note that if this schematic does not contain the kernel arg %q, the machines will not be able to connect to the config server", options.SchematicID, talosConfigKernelArg)
		}
	}

	initScript, err := buildInitScript(options.APIAdvertiseAddress, options.APIPort)
	if err != nil {
		return nil, fmt.Errorf("failed to build init script: %w", err)
	}

	logger.Info("patch iPXE binaries")

	if err = patchBinaries(ctx, initScript, logger); err != nil {
		return nil, err
	}

	logger.Info("successfully patched iPXE binaries")

	kernelArgs := strings.Fields(options.ExtraKernelArgs)

	if configServerEnabled {
		kernelArgs = append(kernelArgs, talosConfigKernelArg)

		logger.Debug("injected talos config kernel arg to the iPXE requests", zap.String("arg", talosConfigKernelArg))
	}

	return &Handler{
		imageFactoryClient: imageFactoryClient,
		options:            options,
		kernelArgs:         kernelArgs,
		initScript:         initScript,
		logger:             logger,
	}, nil
}
