// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipxe

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"go.uber.org/zap"

	"github.com/siderolabs/booter/internal/server/constants"
	"github.com/siderolabs/booter/internal/util"
)

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

// patchBinaries patches iPXE binaries on the fly with the new embedded script.
//
// This relies on special build in `pkgs/ipxe` where a placeholder iPXE script is embedded.
// EFI iPXE binaries are uncompressed, so these are patched directly.
// BIOS amd64 undionly.pxe is compressed, so we instead patch uncompressed version and compress it back using zbin.
// (zbin is built with iPXE).
func patchBinaries(initScript []byte, logger *zap.Logger) error {
	for _, name := range []string{"ipxe", "snp"} {
		if err := patchScript(
			fmt.Sprintf(constants.IPXEPath+"/amd64/%s.efi", name),
			fmt.Sprintf(constants.TFTPPath+"/%s.efi", name),
			initScript,
		); err != nil {
			return fmt.Errorf("failed to patch %q: %w", name, err)
		}

		if err := patchScript(
			fmt.Sprintf(constants.IPXEPath+"/arm64/%s.efi", name),
			fmt.Sprintf(constants.TFTPPath+"/%s-arm64.efi", name),
			initScript,
		); err != nil {
			return fmt.Errorf("failed to patch %q: %w", name, err)
		}
	}

	if err := patchScript(constants.IPXEPath+"/amd64/kpxe/undionly.kpxe.bin", constants.IPXEPath+"/amd64/kpxe/undionly.kpxe.bin.patched", initScript); err != nil {
		return fmt.Errorf("failed to patch undionly.kpxe.bin: %w", err)
	}

	if err := compressKPXE(constants.IPXEPath+"/amd64/kpxe/undionly.kpxe.bin.patched", constants.IPXEPath+"/amd64/kpxe/undionly.kpxe.zinfo",
		constants.TFTPPath+"/undionly.kpxe", logger); err != nil {
		return fmt.Errorf("failed to compress undionly.kpxe: %w", err)
	}

	if err := compressKPXE(constants.IPXEPath+"/amd64/kpxe/undionly.kpxe.bin.patched", constants.IPXEPath+"/amd64/kpxe/undionly.kpxe.zinfo",
		constants.TFTPPath+"/undionly.kpxe.0", logger); err != nil {
		return fmt.Errorf("failed to compress undionly.kpxe.0: %w", err)
	}

	return nil
}

var (
	placeholderStart = []byte("# *PLACEHOLDER START*")
	placeholderEnd   = []byte("# *PLACEHOLDER END*")
)

func patchScript(source, destination string, script []byte) error {
	contents, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	start := bytes.Index(contents, placeholderStart)
	if start == -1 {
		return fmt.Errorf("placeholder start not found in %q", source)
	}

	end := bytes.Index(contents, placeholderEnd)
	if end == -1 {
		return fmt.Errorf("placeholder end not found in %q", source)
	}

	if end < start {
		return fmt.Errorf("placeholder end before start")
	}

	end += len(placeholderEnd)

	length := end - start

	if len(script) > length {
		return fmt.Errorf("script size %d is larger than placeholder space %d", len(script), length)
	}

	script = append(script, bytes.Repeat([]byte{'\n'}, length-len(script))...)

	copy(contents[start:end], script)

	if err = os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}

	return os.WriteFile(destination, contents, 0o644)
}

// compressKPXE is equivalent to: ./util/zbin bin/undionly.kpxe.bin bin/undionly.kpxe.zinfo > bin/undionly.kpxe.zbin.
func compressKPXE(binFile, infoFile, outFile string, logger *zap.Logger) error {
	out, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer util.LogClose(out, logger)

	cmd := exec.Command("/bin/zbin", binFile, infoFile)
	cmd.Stdout = out

	err = cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return fmt.Errorf("zbin failed with exit code %d, stderr: %v", exitErr.ExitCode(), string(exitErr.Stderr))
		}

		return fmt.Errorf("failed to run zbin: %w", err)
	}

	return nil
}
