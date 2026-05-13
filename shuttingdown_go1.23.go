//go:build go1.23
// +build go1.23

package shuttingdown

import (
	"net/http"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var offset_inShutdown = func() uintptr {
	sf, ok := reflect.TypeFor[http.Server]().FieldByName("inShutdown")
	if !ok {
		panic("github.com/bddjr/shuttingdown: failed to get offset of http.Server.inShutdown")
	}
	// Automatic type checking
	if sf.Type != reflect.TypeFor[atomic.Bool]() {
		panic("github.com/bddjr/shuttingdown: failed to check type of http.Server.inShutdown")
	}
	return sf.Offset
}()

// Is [http.Server] shutting down?
func IsShuttingDown(s *http.Server) bool {
	return (*atomic.Bool)(unsafe.Add(unsafe.Pointer(s), offset_inShutdown)).Load()
}
