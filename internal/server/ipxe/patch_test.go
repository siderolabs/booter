// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipxe_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/siderolabs/booter/internal/server/ipxe"
)

// TestPatchBinariesKeySet locks the contract between the served file names and the names the DHCP
// proxy advertises: the flat, arch-suffixed TFTP names and the per-arch HTTP names.
func TestPatchBinariesKeySet(t *testing.T) {
	script := []byte("#!ipxe\nchain http://example.org/boot")

	files, err := ipxe.PatchBinaries(script)
	require.NoError(t, err)

	expected := []string{
		"ipxe.efi", "snp.efi",
		"ipxe-arm64.efi", "snp-arm64.efi",
		"amd64/ipxe.efi", "amd64/snp.efi",
		"arm64/ipxe.efi", "arm64/snp.efi",
		"undionly.kpxe", "undionly.kpxe.0",
	}

	assert.Len(t, files, len(expected))

	for _, name := range expected {
		contents, ok := files[name]

		assert.True(t, ok, "missing file %q", name)
		assert.NotEmpty(t, contents, "file %q is empty", name)
	}

	// the EFI binaries are patched uncompressed, so the script must appear in them verbatim
	assert.True(t, bytes.Contains(files["snp.efi"], script))
}

func TestPatchScript(t *testing.T) {
	placeholder := []byte("head # *PLACEHOLDER START*\n0123456789\n# *PLACEHOLDER END* tail")

	t.Run("success", func(t *testing.T) {
		patched, err := ipxe.PatchScript(bytes.Clone(placeholder), []byte("#!ipxe"))
		require.NoError(t, err)

		assert.True(t, bytes.Contains(patched, []byte("#!ipxe")))
		assert.Len(t, patched, len(placeholder), "patching must preserve the binary size")
	})

	t.Run("placeholder missing", func(t *testing.T) {
		_, err := ipxe.PatchScript([]byte("no placeholder here"), []byte("#!ipxe"))
		assert.ErrorContains(t, err, "placeholder start not found")
	})

	t.Run("end before start", func(t *testing.T) {
		_, err := ipxe.PatchScript([]byte("# *PLACEHOLDER END* # *PLACEHOLDER START*"), []byte("#!ipxe"))
		assert.ErrorContains(t, err, "placeholder end before start")
	})

	t.Run("oversized script", func(t *testing.T) {
		script := bytes.Repeat([]byte{'x'}, len(placeholder)+1)

		_, err := ipxe.PatchScript(bytes.Clone(placeholder), script)
		assert.ErrorContains(t, err, "larger than placeholder space")
	})
}
