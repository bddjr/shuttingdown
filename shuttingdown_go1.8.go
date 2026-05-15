//go:build !go1.11 && go1.8
// +build !go1.11,go1.8

package shuttingdown

import (
	"net/http"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var offset_inShutdown = func() uintptr {
	sf, _ := reflect.TypeOf(http.Server{}).FieldByName("inShutdown")
	return sf.Offset
}()

// Is [http.Server] shutting down?
func IsShuttingDown(s *http.Server) bool {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+offset_inShutdown))) != 0
}
