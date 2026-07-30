// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipxe_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"github.com/siderolabs/booter/internal/server/dhcp"
	"github.com/siderolabs/booter/internal/server/ipxe"
	"github.com/siderolabs/booter/internal/server/server"
)

// TestAdvertisedBootFilesAreServed binds the two halves of the PXE boot contract: every boot file
// name and URL the DHCP proxy hands out must resolve against the patched files booter actually
// serves. The names appear on both sides as independent literals, so this is the test that fails
// when either side drifts.
func TestAdvertisedBootFilesAreServed(t *testing.T) {
	logger := zaptest.NewLogger(t)
	script := []byte("#!ipxe\nchain http://example.org/boot")

	files, err := ipxe.PatchBinaries(script)
	require.NoError(t, err)

	notFound := http.NotFoundHandler()
	handler := server.NewMuxHandler(notFound, notFound, files, logger)

	offer := func(t *testing.T, fwtype dhcp.Firmware) *dhcpv4.DHCPv4 {
		t.Helper()

		req, err := dhcpv4.New(
			dhcpv4.WithMessageType(dhcpv4.MessageTypeRequest),
			dhcpv4.WithHwAddr([]byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}),
		)
		require.NoError(t, err)

		resp, err := dhcp.OfferDHCP(req, "192.168.1.100", 8081, fwtype, dhcp.Port4011)
		require.NoError(t, err)

		return resp
	}

	t.Run("HTTP boot URLs resolve to patched binaries", func(t *testing.T) {
		for _, fwtype := range []dhcp.Firmware{dhcp.FirmwareX86HTTP, dhcp.FirmwareARMHTTP} {
			bootURL, err := url.Parse(offer(t, fwtype).BootFileNameOption())
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, httptest.NewRequestWithContext(t.Context(), http.MethodGet, bootURL.Path, nil))

			require.Equal(t, http.StatusOK, rec.Code, "the advertised URL path %q is not served", bootURL.Path)
			assert.True(t, bytes.Contains(rec.Body.Bytes(), script), "the file served at %q does not contain the patched script", bootURL.Path)
		}
	})

	t.Run("TFTP boot file names are served", func(t *testing.T) {
		// The TFTP server serves the same files map.
		for _, tt := range []struct {
			fwtype       dhcp.Firmware
			checkPatched bool // the BIOS binary is compressed, so the script is not in it verbatim
		}{
			{fwtype: dhcp.FirmwareX86PC},
			{fwtype: dhcp.FirmwareX86EFI, checkPatched: true},
			{fwtype: dhcp.FirmwareARMEFI, checkPatched: true},
		} {
			name := offer(t, tt.fwtype).BootFileName

			contents, ok := files[name]
			require.True(t, ok, "the advertised TFTP boot file %q is not among the served files", name)
			require.NotEmpty(t, contents)

			if tt.checkPatched {
				assert.True(t, bytes.Contains(contents, script), "the advertised TFTP boot file %q does not contain the patched script", name)
			}
		}
	})

	t.Run("iPXE chain URL file is served", func(t *testing.T) {
		bootURL, err := url.Parse(offer(t, dhcp.FirmwareX86Ipxe).BootFileNameOption())
		require.NoError(t, err)

		name := strings.TrimPrefix(bootURL.Path, "/")

		_, ok := files[name]
		require.True(t, ok, "the advertised iPXE chain file %q is not among the served files", name)
	})
}
