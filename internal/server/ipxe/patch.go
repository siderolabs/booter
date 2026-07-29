// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipxe

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/siderolabs/go-zbin/zbin"

	"github.com/siderolabs/booter/internal/server/constants"
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

// bootTemplate is embedded into iPXE binary when that binary is sent to the node.
//
//nolint:dupword,lll
var bootTemplate = template.Must(template.New("iPXE embedded").Parse(`#!ipxe
prompt --key 0x02 --timeout 2000 Press Ctrl-B for the iPXE command line... && shell ||

{{/* print interfaces */}}
ifstat

{{/* retry 10 times overall */}}
set attempts:int32 10
set x:int32 0

:retry_loop

	set idx:int32 0

	:loop
		{{/* try DHCP on each interface */}}
		isset ${net${idx}/mac} || goto exhausted

		ifclose
		iflinkwait --timeout 5000 net${idx} || goto next_iface
		dhcp net${idx} || goto next_iface
		goto boot

	:next_iface
		inc idx && goto loop

	:boot
		{{/* attempt boot, if fails try next iface */}}
		route

		chain --replace http://{{ .Endpoint }}:{{ .Port }}/{{ .ScriptPath }}?uuid=${uuid}&mac=${net${idx}/mac:hexhyp}&domain=${domain}&hostname=${hostname}&serial=${serial}&arch=${buildarch} || goto next_iface

:exhausted
	echo
	echo Failed to iPXE boot successfully via all interfaces

	iseq ${x} ${attempts} && goto fail ||

	echo Retrying...
	echo

	inc x
	goto retry_loop

:fail
	echo
	echo Failed to get a valid response after ${attempts} attempts
	echo

	echo Rebooting in 5 seconds...
	sleep 5
	reboot
`))

func buildInitScript(endpoint string, port int) ([]byte, error) {
	var buf bytes.Buffer

	if err := bootTemplate.Execute(&buf, struct {
		Endpoint   string
		ScriptPath string
		Port       int
	}{
		Endpoint:   endpoint,
		ScriptPath: constants.IPXEURLPath + "/" + bootScriptName,
		Port:       port,
	}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// patchBinaries patches the embedded iPXE binaries with the new embedded script and returns them
// keyed by every name they are served under, over both TFTP and HTTP.
//
// This relies on special build in `pkgs/ipxe` where a placeholder iPXE script is embedded.
// EFI iPXE binaries are uncompressed, so these are patched directly.
// BIOS amd64 undionly.pxe is compressed, so we instead patch uncompressed version and compress it back.
func patchBinaries(initScript []byte) (map[string][]byte, error) {
	files := map[string][]byte{}

	for _, arch := range []struct{ name, suffix string }{
		{name: "amd64", suffix: ""},
		{name: "arm64", suffix: "-arm64"},
	} {
		for _, name := range []string{"ipxe", "snp"} {
			source := "data/" + arch.name + "/" + name + ".efi"

			contents, err := dataFS.ReadFile(source)
			if err != nil {
				return nil, err
			}

			patched, err := patchScript(contents, initScript)
			if err != nil {
				return nil, fmt.Errorf("failed to patch %q: %w", source, err)
			}

			// The flat, architecture-suffixed names are used by the TFTP boot file names, and the
			// per-architecture ones by the HTTP boot URLs handed out by the DHCP proxy.
			files[name+arch.suffix+".efi"] = patched
			files[arch.name+"/"+name+".efi"] = patched
		}
	}

	bin, err := dataFS.ReadFile("data/amd64/kpxe/undionly.kpxe.bin")
	if err != nil {
		return nil, err
	}

	zinfo, err := dataFS.ReadFile("data/amd64/kpxe/undionly.kpxe.zinfo")
	if err != nil {
		return nil, err
	}

	patched, err := patchScript(bin, initScript)
	if err != nil {
		return nil, fmt.Errorf("failed to patch undionly.kpxe.bin: %w", err)
	}

	compressed, err := zbin.Compress(patched, zinfo)
	if err != nil {
		return nil, fmt.Errorf("failed to compress undionly.kpxe: %w", err)
	}

	// The go:embed directive accepts zero-length files, and an empty directive stream compresses to an empty
	// output without an error, which would be served as a "successful" zero-byte boot file.
	if len(compressed) == 0 {
		return nil, fmt.Errorf("compressing undionly.kpxe produced an empty file")
	}

	files["undionly.kpxe"] = compressed
	files["undionly.kpxe.0"] = compressed

	return files, nil
}

var (
	placeholderStart = []byte("# *PLACEHOLDER START*")
	placeholderEnd   = []byte("# *PLACEHOLDER END*")
)

func patchScript(contents, script []byte) ([]byte, error) {
	contents = bytes.Clone(contents) // patch a copy, so the caller's slice is never mutated

	start := bytes.Index(contents, placeholderStart)
	if start == -1 {
		return nil, fmt.Errorf("placeholder start not found")
	}

	end := bytes.Index(contents, placeholderEnd)
	if end == -1 {
		return nil, fmt.Errorf("placeholder end not found")
	}

	if end < start {
		return nil, fmt.Errorf("placeholder end before start")
	}

	end += len(placeholderEnd)

	length := end - start

	if len(script) > length {
		return nil, fmt.Errorf("script size %d is larger than placeholder space %d", len(script), length)
	}

	script = append(bytes.Clone(script), bytes.Repeat([]byte{'\n'}, length-len(script))...)

	copy(contents[start:end], script)

	return contents, nil
}
