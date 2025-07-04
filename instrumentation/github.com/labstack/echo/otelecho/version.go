// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelecho // import "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

// Version is the current release version of the echo instrumentation.
func Version() string {
	return "0.62.0"
	// This string is updated by the pre_release.sh script during release
}
