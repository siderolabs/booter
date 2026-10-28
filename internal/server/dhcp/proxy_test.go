// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package dhcp_test

import (
	"testing"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/iana"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/siderolabs/booter/internal/server/dhcp"
)

func TestOfferDHCP(t *testing.T) {
	t.Parallel()

	const (
		apiAddr = "192.168.1.100"
		apiPort = 50084
	)

	t.Run("port 67 responds with DHCPOFFER", func(t *testing.T) {
		t.Parallel()

		req := newPXEPacket(t, dhcpv4.MessageTypeDiscover)
		resp, err := dhcp.OfferDHCP(req, apiAddr, apiPort, dhcp.FirmwareX86EFI, dhcp.Port67)
		require.NoError(t, err)

		assert.Equal(t, dhcpv4.MessageTypeOffer, resp.MessageType())
		assert.Equal(t, apiAddr, resp.ServerIdentifier().String())
		assert.Equal(t, apiAddr, resp.ServerIPAddr.String())
		assert.Equal(t, "snp.efi", resp.BootFileNameOption())
		assert.Equal(t, "snp.efi", resp.BootFileName)
		assert.NotNil(t, resp.GetOneOption(dhcpv4.OptionClassIdentifier))
	})

	t.Run("port 4011 responds with DHCPACK", func(t *testing.T) {
		t.Parallel()

		req := newPXEPacket(t, dhcpv4.MessageTypeRequest)
		resp, err := dhcp.OfferDHCP(req, apiAddr, apiPort, dhcp.FirmwareX86EFI, dhcp.Port4011)
		require.NoError(t, err)

		assert.Equal(t, dhcpv4.MessageTypeAck, resp.MessageType())
		assert.Equal(t, apiAddr, resp.ServerIdentifier().String())
		assert.Equal(t, apiAddr, resp.ServerIPAddr.String())
		assert.Equal(t, "snp.efi", resp.BootFileNameOption())
		assert.Equal(t, "snp.efi", resp.BootFileName)
		assert.NotNil(t, resp.GetOneOption(dhcpv4.OptionClassIdentifier))
	})

	t.Run("firmware types produce correct boot filenames", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			wantFile      string
			wantBOOTPFile string // BOOTP `file` header (resp.BootFileName); empty when only DHCP option 67 is set.
			fwtype        dhcp.Firmware
		}{
			{fwtype: dhcp.FirmwareX86PC, wantFile: "undionly.kpxe", wantBOOTPFile: "undionly.kpxe"},
			{fwtype: dhcp.FirmwareX86EFI, wantFile: "snp.efi", wantBOOTPFile: "snp.efi"},
			{fwtype: dhcp.FirmwareARMEFI, wantFile: "snp-arm64.efi", wantBOOTPFile: "snp-arm64.efi"},
			{fwtype: dhcp.FirmwareX86Ipxe, wantFile: "tftp://192.168.1.100/undionly.kpxe"},
			{fwtype: dhcp.FirmwareX86HTTP, wantFile: "http://192.168.1.100:50084/tftp/amd64/snp.efi"},
			{fwtype: dhcp.FirmwareARMHTTP, wantFile: "http://192.168.1.100:50084/tftp/arm64/snp.efi"},
		}

		for _, tt := range tests {
			req := newPXEPacket(t, dhcpv4.MessageTypeDiscover)
			resp, err := dhcp.OfferDHCP(req, apiAddr, apiPort, tt.fwtype, dhcp.Port67)
			require.NoError(t, err)

			assert.Equal(t, apiAddr, resp.ServerIdentifier().String(), "firmware type %d option 54", tt.fwtype)
			assert.Equal(t, tt.wantFile, resp.BootFileNameOption(), "firmware type %d option 67", tt.fwtype)
			assert.Equal(t, tt.wantBOOTPFile, resp.BootFileName, "firmware type %d BOOTP file header", tt.fwtype)
		}
	})
}

func newPXEPacket(t *testing.T, msgType dhcpv4.MessageType) *dhcpv4.DHCPv4 {
	t.Helper()

	pkt, err := dhcpv4.New(
		dhcpv4.WithMessageType(msgType),
		dhcpv4.WithHwAddr([]byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}),
	)
	require.NoError(t, err)

	// Option 93: Client System Architecture (EFI x86_64)
	pkt.UpdateOption(dhcpv4.OptClientArch(iana.EFI_X86_64))

	return pkt
}
