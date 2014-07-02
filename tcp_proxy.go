package peroxide

import (
	"io"
	"net"
)

type tcpConn struct {
	t        TestingT
	src, dst net.Conn
}

func newTCPConn(t TestingT, src, dst net.Conn) *tcpConn {
	cn := &tcpConn{t: t, src: src, dst: dst}

	go cn.proxy(cn.src, cn.dst)
	go cn.proxy(cn.dst, cn.src)

	return cn
}

func (cn *tcpConn) proxy(dst, src net.Conn) {
	_, err := io.Copy(dst, src)
	if err == nil {
		dst.Close()
	}
}

func (cn *tcpConn) Close() {
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

func (l *TCPListener) AcceptOne() (net.Addr, <-chan Conn) {
	dst, err := net.Dial("tcp", l.dstAddr)
	if err != nil {
		l.t.Fatal(err)
	}

	ln, err := net.Listen("tcp", l.listenAddr)
	if err != nil {
		l.t.Fatal(err)
	}

	ch := make(chan Conn, 1)
	go func() {

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

	return ln.Addr(), ch
}

func (l *TCPListener) Close() {
}
