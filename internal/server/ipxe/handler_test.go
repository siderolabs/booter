// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipxe_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"github.com/siderolabs/booter/internal/server/ipxe"
)

// TestNewHandlerPatchesEmbeddedBinaries checks that the iPXE binaries booter embeds, from the iPXE image pinned
// in its .kres.yaml, can be patched by the shared iPXE package, which is tested against the provider's binaries.
func TestNewHandlerPatchesEmbeddedBinaries(t *testing.T) {
	t.Parallel()

	handler, err := ipxe.NewHandler(false, nil, ipxe.HandlerOptions{
		APIAdvertiseAddress: "192.168.1.100",
		APIPort:             8081,
		TalosVersion:        "v1.13.0",
	}, zaptest.NewLogger(t))
	require.NoError(t, err)

	require.NotEmpty(t, handler.PatchedFiles()["snp.efi"])
	require.NotEmpty(t, handler.PatchedFiles()["undionly.kpxe"])
}
