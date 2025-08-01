// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package omni

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cosi-project/runtime/pkg/safe"
	"github.com/siderolabs/omni/client/pkg/client"
	"github.com/siderolabs/omni/client/pkg/omni/resources/siderolink"
	"go.uber.org/zap"
)

// Options defines the options for building the machine configuration.
type Options struct {
	APIEndpoint              string
	APIInsecureSkipTLSVerify bool
}

type ConnectionOptions struct {
	SiderolinkAPIURL string
	JoinToken        string

	EventsPort  int
	KmsgLogPort int
}

func GetConnectionOptions(ctx context.Context, apiEndpoint string, insecureSkipTLSVerify bool, logger *zap.Logger) (ConnectionOptions, error) {
	omniClient, err := buildAPIClient(apiEndpoint, insecureSkipTLSVerify)
	if err != nil {
		return ConnectionOptions{}, fmt.Errorf("failed to build Omni API client: %w", err)
	}

	st := omniClient.Omni().State()

	connectionParams, err := safe.StateGetByID[*siderolink.ConnectionParams](ctx, st, siderolink.ConfigID)
	if err != nil {
		return ConnectionOptions{}, fmt.Errorf("failed to get Omni connection params: %w", err)
	}

	logger.Info("fetched Omni connection parameters",
		zap.String("api_endpoint", connectionParams.TypedSpec().Value.ApiEndpoint),
		zap.Int("events_port", int(connectionParams.TypedSpec().Value.EventsPort)),
		zap.Int("kmsg_log_port", int(connectionParams.TypedSpec().Value.LogsPort)),
	)

	return ConnectionOptions{
		SiderolinkAPIURL: connectionParams.TypedSpec().Value.ApiEndpoint,
		JoinToken:        connectionParams.TypedSpec().Value.JoinToken,
		EventsPort:       int(connectionParams.TypedSpec().Value.EventsPort),
		KmsgLogPort:      int(connectionParams.TypedSpec().Value.LogsPort),
	}, nil
}

// buildOmniAPIClient creates a new Omni API client.
func buildAPIClient(apiEndpoint string, insecureSkipTLSVerify bool) (*client.Client, error) {
	if apiEndpoint == "" {
		return nil, errors.New("omni api endpoint is not set")
	}

	const omniServiceAccountKeyEnvVar = "OMNI_SERVICE_ACCOUNT_KEY"

	serviceAccountKey := os.Getenv(omniServiceAccountKeyEnvVar)

	if serviceAccountKey == "" {
		return nil, fmt.Errorf("environment variable %q is not set", omniServiceAccountKeyEnvVar)
	}

	return client.New(apiEndpoint, client.WithServiceAccount(serviceAccountKey), client.WithInsecureSkipTLSVerify(insecureSkipTLSVerify))
}
