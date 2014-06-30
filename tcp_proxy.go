package peroxide

import (
	"io"
	"net"
)

// TCPConn implements the Conn interface for TCP connections.
type TCPConn struct {
	t        TestingT
	src, dst net.Conn
}

func newTCPConn(t TestingT, src, dst net.Conn) *TCPConn {
	cn := &TCPConn{t: t, src: src, dst: dst}

	go cn.proxy(cn.src, cn.dst)
	go cn.proxy(cn.dst, cn.src)

	return cn
}

func (cn *TCPConn) proxy(dst, src net.Conn) {
	_, err := io.Copy(dst, src)
	if err == nil {
		dst.Close()
	}
}

func (cn *TCPConn) Close() {
	cn.src.Close()
	cn.dst.Close()
}

// TCPListener implements the Listener interface for the TCP protocol.
type TCPListener struct {
	t                   TestingT
	listenAddr, dstAddr string
}

func NewTCPListener(t TestingT, listenAddr, dstAddr string) *TCPListener {
	return &TCPListener{t: t, listenAddr: listenAddr, dstAddr: dstAddr}
}

func (l *TCPListener) AcceptOne() <-chan *TCPConn {
	dst, err := net.Dial("tcp", l.dstAddr)
	if err != nil {
		l.t.Fatal(err)
	}

	ch := make(chan *TCPConn)
	go func() {
		ln, err := net.Listen("tcp", l.listenAddr)
		if err != nil {
			l.t.Error(err)
		}

		src, err := ln.Accept()
		if err != nil {
			l.t.Error(err)
		}

		err = ln.Close()
		if err != nil {
			l.t.Error(err)
		}

		ch <- newTCPConn(l.t, src, dst)
		close(ch)
	}()
	return ch
}

func (l *TCPListener) Close() {
}
