/*
Package peroxide provides some simple proxy-server-like functions to simplify testing reliable or redundant network services and bindings.

These projects have corner cases that are otherwise difficult to test as they involve bringing the service up/down at specific intervals in order
to test the correct availability sequence. Peroxide makes this somewhat easier to accomplish, as you don't have to touch the running service; simply
start/stop a proxy at will.
*/
package peroxide

// TestingT is an interface wrapping golang's built-in *testing.T.
type TestingT interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// Conn is an interface implemented by all proxied connections.
type Conn interface {
	Close()
}

// Listener is an interface for connection-oriented protocols like TCP.
type Listener interface {
	AcceptOne() <-chan *Conn
	Close()
}
