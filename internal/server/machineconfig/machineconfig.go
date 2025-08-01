// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package machineconfig builds the machine configuration for the server.
package machineconfig

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/siderolabs/talos/pkg/machinery/config/config"
	"github.com/siderolabs/talos/pkg/machinery/config/container"
	"github.com/siderolabs/talos/pkg/machinery/config/encoder"
	"github.com/siderolabs/talos/pkg/machinery/config/types/meta"
	"github.com/siderolabs/talos/pkg/machinery/config/types/runtime"
	"github.com/siderolabs/talos/pkg/machinery/config/types/siderolink"

	"github.com/siderolabs/booter/internal/server/omni"
)

const siderolinkAddress = "fdae:41e4:649b:9303::1"

// Build builds the machine configuration for the server.
func Build(omniConnOpts omni.ConnectionOptions) ([]byte, error) {
	apiURL, err := url.Parse(omniConnOpts.SiderolinkAPIURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API URL: %w", err)
	}

	siderolinkConfig := siderolink.NewConfigV1Alpha1()
	siderolinkConfig.APIUrlConfig = meta.URL{
		URL: apiURL,
	}

	eventSinkConfig := runtime.NewEventSinkV1Alpha1()
	eventSinkConfig.Endpoint = net.JoinHostPort(siderolinkAddress, strconv.Itoa(omniConnOpts.EventsPort))

	kmsgLogURL, err := url.Parse("tcp://" + net.JoinHostPort(siderolinkAddress, strconv.Itoa(omniConnOpts.KmsgLogPort)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse kmsg log URL: %w", err)
	}

	kmsgLogConfig := runtime.NewKmsgLogV1Alpha1()
	kmsgLogConfig.MetaName = "omni-kmsg"
	kmsgLogConfig.KmsgLogURL = meta.URL{
		URL: kmsgLogURL,
	}

	documents := []config.Document{
		siderolinkConfig,
		eventSinkConfig,
		kmsgLogConfig,
	}

	configContainer, err := container.New(documents...)
	if err != nil {
		return nil, fmt.Errorf("failed to create config container: %w", err)
	}

	return configContainer.EncodeBytes(encoder.WithComments(encoder.CommentsDisabled))
}
