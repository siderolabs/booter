// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package main implements the main entrypoint for the server.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/siderolabs/booter/internal/server"
	"github.com/siderolabs/booter/internal/version"
)

const (
	extraKernelArgsEnvVar = "EXTRA_KERNEL_ARGS"
	extraKernelArgsFlag   = "extra-kernel-args"
)

var (
	serverOptions = server.DefaultOptions()
	debug         bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     version.Name,
	Short:   "Run the server",
	Version: version.Tag,
	Args:    cobra.ArbitraryArgs,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		cmd.SilenceUsage = true // if the args are parsed fine, no need to show usage
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, err := initLogger()
		if err != nil {
			return fmt.Errorf("failed to create logger: %w", err)
		}

		defer logger.Sync() //nolint:errcheck

		return run(cmd.Context(), args, logger)
	},
}

func initLogger() (*zap.Logger, error) {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if !debug {
		loggerConfig.Level.SetLevel(zap.InfoLevel)
	}

	return loggerConfig.Build(zap.AddStacktrace(zapcore.FatalLevel)) // only print stack traces for fatal errors)
}

func run(ctx context.Context, args []string, logger *zap.Logger) error {
	serverOptions.ExtraKernelArgs = getExtraKernelArgs(args, logger)

	prov := server.New(serverOptions, logger)

	if err := prov.Run(ctx); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}

func getExtraKernelArgs(args []string, logger *zap.Logger) string {
	if serverOptions.ExtraKernelArgs != "" {
		logger.Info(fmt.Sprintf("use extra kernel args from %q flag", "--"+extraKernelArgsFlag), zap.String("args", serverOptions.ExtraKernelArgs))

		return serverOptions.ExtraKernelArgs
	}

	kernelArgsEnv := os.Getenv(extraKernelArgsEnvVar)
	if kernelArgsEnv != "" {
		logger.Info(fmt.Sprintf("use extra kernel args from %q environment variable", extraKernelArgsEnvVar), zap.String("args", kernelArgsEnv))

		return kernelArgsEnv
	}

	if len(args) > 0 {
		joined := strings.Join(args, " ")

		logger.Info("use extra kernel args from command line arguments", zap.String("args", joined))

		return joined
	}

	return ""
}

func main() {
	if err := runCmd(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func runCmd() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode & logs.")

	rootCmd.Flags().StringVar(&serverOptions.APIListenAddress, "api-listen-address", serverOptions.APIListenAddress,
		"The IP address to listen on. If not specified, the server will listen on all interfaces.")
	rootCmd.Flags().StringVar(&serverOptions.APIAdvertiseAddress, "api-advertise-address", serverOptions.APIAdvertiseAddress,
		"The IP address to advertise. Required if the server has more than a single routable IP address. If not specified, the single routable IP address will be used.")
	rootCmd.Flags().IntVar(&serverOptions.APIPort, "api-port", serverOptions.APIPort, "The port to run the api server on.")
	rootCmd.Flags().StringVar(&serverOptions.DHCPProxyIfaceOrIP, "dhcp-proxy-iface-or-ip", serverOptions.DHCPProxyIfaceOrIP,
		"The interface name or the IP address on the interface to run the DHCP proxy server on. "+
			"If it is an IP address, the DHCP proxy server will run on the interface that has the IP address. "+
			"If not specified, defaults to the API advertise address.")
	rootCmd.Flags().StringVar(&serverOptions.ImageFactoryBaseURL, "image-factory-base-url", serverOptions.ImageFactoryBaseURL,
		"The base URL of the image factory.")
	rootCmd.Flags().StringVar(&serverOptions.ImageFactoryPXEBaseURL, "image-factory-pxe-base-url", serverOptions.ImageFactoryPXEBaseURL,
		"The base URL of the image factory PXE server.")
	rootCmd.Flags().BoolVar(&serverOptions.SecureBootEnabled, "secure-boot-enabled", serverOptions.SecureBootEnabled,
		"Serve secure boot UKI from the iPXE endpoint. The UKI can be used to boot a machine without secure boot, but it is required to boot a machine with secure boot.",
	)

	rootCmd.Flags().BoolVar(&serverOptions.DisableDHCPProxy, "disable-dhcp-proxy", serverOptions.DisableDHCPProxy,
		"Disable the DHCP proxy server.")

	rootCmd.Flags().StringVar(&serverOptions.SchematicID, "schematic-id", serverOptions.SchematicID,
		"The ID of the schematic to use.")
	rootCmd.Flags().StringSliceVar(&serverOptions.Extensions, "extensions", serverOptions.Extensions,
		"List of Talos extensions to use. The extensions will be used to generate schematic ID from the image factory.")
	rootCmd.Flags().StringVarP(&serverOptions.ExtraKernelArgs, extraKernelArgsFlag, "k", serverOptions.ExtraKernelArgs,
		fmt.Sprintf("List of extra kernel arguments to use. "+
			"They can be used, e.g., to connect the machines to Omni over SideroLink."+
			"The arguments will be used to generate schematic ID from the image factory. "+
			"These extra args can also be set via the %q environment variable or via command line arguments.", extraKernelArgsEnvVar))

	rootCmd.Flags().StringVar(&serverOptions.TalosVersion, "talos-version", serverOptions.TalosVersion, "The Talos version to use.")

	// Omni options
	// todo: disabled for now, we can re-enable it after https://github.com/siderolabs/omni/issues/1375 for better UX
	// rootCmd.Flags().StringVar(&serverOptions.Omni.APIEndpoint, "omni-api-endpoint", serverOptions.Omni.APIEndpoint,
	//	"The endpoint of the Omni API. If specified, Talos machines will be connected to Omni.")
	// rootCmd.Flags().BoolVar(&serverOptions.Omni.APIInsecureSkipTLSVerify, "omni-api-insecure-skip-tls-verify", serverOptions.Omni.APIInsecureSkipTLSVerify,
	//	"Skip TLS verification for the Omni API endpoint. This is useful for development purposes, but should not be used in production.")
}
