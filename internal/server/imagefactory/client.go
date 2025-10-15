// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package imagefactory

import (
	"context"
	"fmt"
	"time"

	"github.com/blang/semver/v4"
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

// EnsureSchematic ensures a schematic exists on the image factory and returns its ID.
func (c *Client) EnsureSchematic(ctx context.Context, extensions, extraKernelArgs []string) (string, error) {
	logger := c.logger.With(zap.Strings("extensions", extensions), zap.Strings("extra_kernel_args", extraKernelArgs))

	logger.Debug("ensure schematic")

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

	return schematicID, nil
}

// GetIPXEURL returns the iPXE URL for the given schematic ID, Talos version, and architecture.
func (c *Client) GetIPXEURL(schematicID, talosVersion, arch string) (string, error) {
	if schematicID == "" {
		return "", fmt.Errorf("schematic ID is required")
	}

	if talosVersion == "" {
		return "", fmt.Errorf("talos version is required")
	}

	if arch == "" {
		return "", fmt.Errorf("arch is required")
	}

	ipxeURL := fmt.Sprintf("%s/pxe/%s/%s/metal-%s", c.pxeBaseURL, schematicID, talosVersion, arch)

	if c.secureBootEnabled {
		ipxeURL += "-secureboot"
	}

	return ipxeURL, nil
}

// GetLatestStableVersion returns the latest stable Talos version from the image factory.
func (c *Client) GetLatestStableVersion(ctx context.Context) (string, error) {
	versions, err := c.factoryClient.Versions(ctx)
	if err != nil {
		return "", err
	}

	var latestStable *semver.Version
	for _, v := range versions {
		sv, err := semver.ParseTolerant(v)
		if err != nil {
			c.logger.Warn("failed to parse version", zap.String("version", v), zap.Error(err))

			continue
		}

		// Skip pre-releases
		if len(sv.Pre) > 0 {
			continue
		}

		if latestStable == nil || sv.GT(*latestStable) {
			latestStable = &sv
		}
	}

	if latestStable == nil {
		return "", fmt.Errorf("no stable versions found")
	}

	return latestStable.String(), nil
}
