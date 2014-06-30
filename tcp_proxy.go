package peroxide

import (
	"io"
	"net"
)

type TCPConn struct {
	t        TestingT
	src, dst net.Conn
}

func newTCPConn(t TestingT, src net.Conn, dstAddr string) *TCPConn {
	cn := &TCPConn{t: t, src: src}

	dst, err := net.Dial("tcp", dstAddr)
	if err != nil {
		t.Fatal("proxy couldn't connect", err)
	}

	cn.dst = dst

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

type TCPListener struct {
	t                   TestingT
	listenAddr, dstAddr string
}

func NewTCPListener(t TestingT, listenAddr, dstAddr string) *TCPListener {
	return &TCPListener{t: t, listenAddr: listenAddr, dstAddr: dstAddr}
}

func (l *TCPListener) AcceptOne() <-chan *TCPConn {
	ch := make(chan *TCPConn)
	go func() {
		ln, err := net.Listen("tcp", l.listenAddr)
		if err != nil {
			l.t.Fatal(err)
		}

		src, err := ln.Accept()
		if err != nil {
			l.t.Fatal(err)
		}

		err = ln.Close()
		if err != nil {
			l.t.Fatal(err)
		}

		ch <- newTCPConn(l.t, src, l.dstAddr)
		close(ch)
	}()
	return ch
}

func (l *TCPListener) Close() {
}
