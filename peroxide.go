package peroxide

type TestingT interface {
	Error(args ...interface{})
}

type Conn interface {
	Close()
}

type Listener interface {
	AcceptOne() <-chan *Conn
	Close()
}
