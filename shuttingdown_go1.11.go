//go:build !go1.23 && go1.11
// +build !go1.23,go1.11

package shuttingdown

import (
	"net/http"
	_ "unsafe"
)

// Is [http.Server] shutting down?
//
//go:linkname IsShuttingDown net/http.(*Server).shuttingDown
func IsShuttingDown(*http.Server) bool
