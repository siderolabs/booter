// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipxe

import (
	"embed"
	"io/fs"
)

// dataFS holds the iPXE binaries embedded into the booter binary.
//
// The docker build populates the data directory from the iPXE image stages, and a native build
// populates it with "make fetch-source-assets". The files are listed explicitly, so a binary can
// never be built without them.
//
//go:embed data/amd64/ipxe.efi data/amd64/snp.efi
//go:embed data/amd64/kpxe/undionly.kpxe.bin data/amd64/kpxe/undionly.kpxe.zinfo
//go:embed data/arm64/ipxe.efi data/arm64/snp.efi
var dataFS embed.FS

// binaries returns the embedded iPXE binaries, laid out as the iPXE package's PatchBinaries expects.
func binaries() fs.FS {
	binaries, err := fs.Sub(dataFS, "data")
	if err != nil { // only fails on an invalid path, which "data" is not
		panic(err)
	}

	return binaries
}
