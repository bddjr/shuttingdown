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
	const errmsg = "github.com/bddjr/shuttingdown: failed to check type of http.Server.inShutdown"
	b := reflect.TypeFor[atomic.Bool]()
	if sf.Type.Kind() != b.Kind() {
		panic(errmsg)
	}
	if sf.Type.PkgPath() != b.PkgPath() {
		panic(errmsg)
	}
	if sf.Type.Name() != b.Name() {
		panic(errmsg)
	}
	return sf.Offset
}()

// Is [http.Server] shutting down?
func IsShuttingDown(s *http.Server) bool {
	return (*atomic.Bool)(unsafe.Add(unsafe.Pointer(s), offset_inShutdown)).Load()
}
