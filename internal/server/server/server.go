// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package server implements the HTTP and GRPC servers.
package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/siderolabs/booter/internal/server/constants"
)

// Server represents the HTTP server.
type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

// New creates a new server.
func New(ctx context.Context, listenAddress string, port int, configHandler, ipxeHandler http.Handler, files map[string][]byte, logger *zap.Logger) *Server {
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(listenAddress, strconv.Itoa(port)),
		Handler: NewMuxHandler(configHandler, ipxeHandler, files, logger),
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	return &Server{
		httpServer: httpServer,
		logger:     logger,
	}
}

// Run runs the server.
func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.shutdownOnCancel(ctx, s.httpServer)
	})

	eg.Go(func() error {
		s.logger.Info("start HTTP server", zap.String("address", s.httpServer.Addr))

		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to run server: %w", err)
		}

		return nil
	})

	return eg.Wait()
}

func (s *Server) shutdownOnCancel(ctx context.Context, server *http.Server) error {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil { //nolint:contextcheck
		return fmt.Errorf("failed to shutdown HTTP server: %w", err)
	}

	return nil
}

// NewMuxHandler creates the HTTP handler that multiplexes the config, the iPXE scripts, and the
// boot file serving.
func NewMuxHandler(configHandler, ipxeHandler http.Handler, files map[string][]byte, logger *zap.Logger) http.Handler {
	mux := http.NewServeMux()

	if configHandler != nil {
		mux.Handle("/"+constants.ConfigURLPath, configHandler)
	}

	mux.Handle(fmt.Sprintf("/%s/{script}", constants.IPXEURLPath), ipxeHandler)
	mux.Handle("/tftp/", http.StripPrefix("/tftp/", filesHandler(files)))

	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, req)

			logger.Info(
				"request",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}

	return loggingMiddleware(mux)
}

// filesHandler serves the given in-memory files, keyed by the request path.
func filesHandler(files map[string][]byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contents, ok := files[strings.TrimPrefix(path.Clean("/"+req.URL.Path), "/")]
		if !ok {
			http.NotFound(w, req)

			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")

		// ServeContent handles HEAD, range, and conditional requests, and sets Content-Length.
		http.ServeContent(w, req, "", time.Time{}, bytes.NewReader(contents))
	})
}
