// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package constants provides constants for the server package.
package constants

const (
	// IPXEPath is the path to the iPXE binaries.
	IPXEPath = "/var/lib/ipxe"

	// TFTPPath is the path from which the TFTP server serves files.
	TFTPPath = "/var/lib/tftp"

	// IPXEURLPath is the path from which the HTTP server serves the iPXE scripts.
	IPXEURLPath = "ipxe"
)
