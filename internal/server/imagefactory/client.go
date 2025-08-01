// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package imagefactory

import (
	"context"
	"fmt"
	"time"

	"github.com/siderolabs/image-factory/pkg/client"
	"github.com/siderolabs/image-factory/pkg/schematic"
	"go.uber.org/zap"
)

// Client is an image factory client.
type Client struct {
	factoryClient     *client.Client
	logger            *zap.Logger
	pxeBaseURL        string
	secureBootEnabled bool
}

// NewClient creates a new image factory client.
func NewClient(baseURL, pxeBaseURL string, secureBootEnabled bool, logger *zap.Logger) (*Client, error) {
	factoryClient, err := client.New(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		pxeBaseURL:        pxeBaseURL,
		factoryClient:     factoryClient,
		secureBootEnabled: secureBootEnabled,
		logger:            logger,
	}, nil
}

// SchematicIPXEURL ensures a schematic exists on the image factory and returns the iPXE URL to it.
//
// If agentMode is true, the schematic will be created with the firmware extensions and the metal-agent extension.
func (c *Client) SchematicIPXEURL(ctx context.Context, talosVersion, arch string, extensions, extraKernelArgs []string) (string, error) {
	logger := c.logger.With(zap.String("talos_version", talosVersion), zap.String("arch", arch),
		zap.Strings("extensions", extensions), zap.Strings("extra_kernel_args", extraKernelArgs))

	logger.Debug("generate schematic iPXE URL")

	if talosVersion == "" {
		return "", fmt.Errorf("talosVersion is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	sch := schematic.Schematic{
		Customization: schematic.Customization{
			ExtraKernelArgs: extraKernelArgs,
			SystemExtensions: schematic.SystemExtensions{
				OfficialExtensions: extensions,
			},
		},
	}

	marshaled, err := sch.Marshal()
	if err != nil {
		return "", fmt.Errorf("failed to marshal schematic: %w", err)
	}

	logger.Debug("generated schematic", zap.String("schematic", string(marshaled)))

	schematicID, err := c.factoryClient.SchematicCreate(ctx, sch)
	if err != nil {
		return "", fmt.Errorf("failed to create schematic: %w", err)
	}

	ipxeURL := fmt.Sprintf("%s/pxe/%s/%s/metal-%s", c.pxeBaseURL, schematicID, talosVersion, arch)
	if c.secureBootEnabled {
		ipxeURL += "-secureboot"
	}

	logger.Debug("generated schematic iPXE URL", zap.String("ipxe_url", ipxeURL))

	return ipxeURL, nil
}
