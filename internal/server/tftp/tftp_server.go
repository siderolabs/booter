// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package tftp implements a TFTP server.
package tftp

import (
	"context"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/pin/tftp/v3"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/siderolabs/booter/internal/server/constants"
	"github.com/siderolabs/booter/internal/util"
)

// Server represents the TFTP server serving iPXE binaries.
type Server struct {
	logger *zap.Logger

	listenAddress string
}

// NewServer creates a new TFTP server.
func NewServer(listenAddress string, logger *zap.Logger) *Server {
	return &Server{
		listenAddress: listenAddress,
		logger:        logger,
	}
}

// Run runs the TFTP server.
func (s *Server) Run(ctx context.Context) error {
	if err := os.MkdirAll(constants.TFTPPath, 0o777); err != nil {
		return err
	}

	readHandler := func(filename string, rf io.ReaderFrom) error {
		return handleRead(filename, rf, s.logger)
	}

	srv := tftp.NewServer(readHandler, nil)

	// A standard TFTP server implementation receives requests on port 69 and
	// allocates a new high port (over 1024) dedicated to that request. In single
	// port mode, the same port is used for transmit and receive. If the server
	// is started on port 69, all communication will be done on port 69.
	// This option is required since the Kubernetes service definition defines a
	// single port.
	srv.EnableSinglePort()
	srv.SetTimeout(5 * time.Second)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return srv.ListenAndServe(net.JoinHostPort(s.listenAddress, "69"))
	})

	eg.Go(func() error {
		<-ctx.Done()

		srv.Shutdown()

		return nil
	})

	return eg.Wait()
}

// cleanPath makes a path safe for use with filepath.Join. This is done by not
// only cleaning the path, but also (if the path is relative) adding a leading
// '/' and cleaning it (then removing the leading '/'). This ensures that a
// path resulting from prepending another path will always resolve to lexically
// be a subdirectory of the prefixed path. This is all done lexically, so paths
// that include symlinks won't be safe as a result of using CleanPath.
func cleanPath(path string) string {
	// Deal with empty strings nicely.
	if path == "" {
		return ""
	}

	// Ensure that all paths are cleaned (especially problematic ones like
	// "/../../../../../" which can cause lots of issues).
	path = filepath.Clean(path)

	// If the path isn't absolute, we need to do more processing to fix paths
	// such as "../../../../<etc>/some/path". We also shouldn't convert absolute
	// paths to relative ones.
	if !filepath.IsAbs(path) {
		path = filepath.Clean(string(os.PathSeparator) + path)
		// This can't fail, as (by definition) all paths are relative to root.
		path, _ = filepath.Rel(string(os.PathSeparator), path) //nolint:errcheck
	}

	// Clean the path again for good measure.
	return filepath.Clean(path)
}

// handleRead is called when a client starts file download from server.
func handleRead(filename string, rf io.ReaderFrom, logger *zap.Logger) error {
	logger.Info("file requested", zap.String("filename", filename))

	filename = filepath.Join(constants.TFTPPath, cleanPath(filename))

	file, err := os.Open(filename)
	if err != nil {
		logger.Error("failed to open file", zap.String("filename", filename), zap.Error(err))

		return err
	}

	defer util.LogClose(file, logger)

	n, err := rf.ReadFrom(file)
	if err != nil {
		logger.Error("failed to read from file", zap.String("filename", filename), zap.Error(err))

		return err
	}

	logger.Info("file sent", zap.String("filename", filename), zap.Int64("bytes", n))

	return nil
}
