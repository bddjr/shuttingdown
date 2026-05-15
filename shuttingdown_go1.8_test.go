//go:build !go1.11 && go1.8
// +build !go1.11,go1.8

package shuttingdown_test

import (
	"context"
	"errors"
	"net"
	"net/http"
	"testing"

	"github.com/bddjr/shuttingdown"
)

type listener struct {
	s        *http.Server
	close    chan struct{}
	callback chan struct{}
}

func (l *listener) Accept() (net.Conn, error) {
	l.s.Shutdown(context.Background())
	<-l.close
	return nil, errors.New("listener closed")
}

func (l *listener) Close() error {
	select {
	case <-l.close:
		//
	default:
		close(l.close)
		<-l.callback
	}
	return nil
}

func (l *listener) Addr() net.Addr {
	return nil
}

func Test(t *testing.T) {
	s := &http.Server{}
	if shuttingdown.IsShuttingDown(s) {
		panic(1)
	}

	l := listener{
		s:        s,
		close:    make(chan struct{}),
		callback: make(chan struct{}),
	}
	go s.Serve(&l)
	<-l.close
	defer close(l.callback)

	if !shuttingdown.IsShuttingDown(s) {
		panic(2)
	}
	println("ok")
}
