// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package config serves machine configuration to the machines that request it via talos.config kernel argument.
package config

import (
	"net/http"

	"go.uber.org/zap"
)

// Handler handles machine configuration requests.
type Handler struct {
	logger        *zap.Logger
	machineConfig []byte
}

// NewHandler creates a new Handler.
func NewHandler(machineConfig []byte, logger *zap.Logger) (*Handler, error) {
	return &Handler{
		machineConfig: machineConfig,
		logger:        logger,
	}, nil
}

// ServeHTTP serves the machine configuration.
//
// URL pattern: http://ip-of-this-server:50084/config?&u=${uuid}
//
// Implements http.Handler interface.
func (s *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("u")

	s.logger.Info("handle config request", zap.String("uuid", uuid))

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(s.machineConfig); err != nil {
		s.logger.Error("failed to write response", zap.Error(err))
	}
}
