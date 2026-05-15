//go:build go1.11
// +build go1.11

package shuttingdown_test

import (
	"net/http"
	"testing"

	"github.com/bddjr/shuttingdown"
)

func Test(t *testing.T) {
	s := &http.Server{}
	if shuttingdown.IsShuttingDown(s) {
		panic(1)
	}
	s.Close()
	if !shuttingdown.IsShuttingDown(s) {
		panic(2)
	}
	println("ok")
}
