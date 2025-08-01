// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package util provides utility functions.
package util

import (
	"io"

	"go.uber.org/zap"
)

// LogClose closes the closer and logs any error that occurs.
func LogClose(closer io.Closer, logger *zap.Logger) {
	if err := closer.Close(); err != nil {
		logger.Error("failed to close", zap.Error(err))
	}
}
