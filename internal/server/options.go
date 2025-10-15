// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import "github.com/siderolabs/booter/internal/server/omni"

// Options contains the server options.
type Options struct {
	ImageFactoryBaseURL    string
	ImageFactoryPXEBaseURL string
	APIListenAddress       string
	APIAdvertiseAddress    string
	DHCPProxyIfaceOrIP     string
	TalosVersion           string
	ExtraKernelArgs        string
	SchematicID            string
	Extensions             []string
	Omni                   omni.Options
	APIPort                int
	DisableDHCPProxy       bool
	SecureBootEnabled      bool
}

// DefaultOptions returns the default server options.
func DefaultOptions() Options {
	return Options{
		ImageFactoryBaseURL:    "https://factory.talos.dev",
		ImageFactoryPXEBaseURL: "https://pxe.factory.talos.dev",
		APIPort:                50084,
	}
}
