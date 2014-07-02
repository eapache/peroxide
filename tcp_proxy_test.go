package peroxide

import (
	"io/ioutil"
	"net"
	"testing"
)

func TestTCPListener(t *testing.T) {
	ln, err := net.Listen("tcp", ":")
	if err != nil {
		t.Error(err)
	}

	done := make(chan struct{})
	go func() {
		src, err := ln.Accept()
		if err != nil {
			t.Error(err)
		}

		err = ln.Close()
		if err != nil {
			t.Error(err)
		}

		val, err := ioutil.ReadAll(src)
		if err != nil {
			t.Error(err)
		}

		if string(val) != "TESTING 1 2 3" {
			t.Error("failed to read correct val:", string(val))
		}

		close(done)
	}()

	var l Listener
	l = NewTCPListener(t, ":", ln.Addr().String())
	addr, _ := l.AcceptOne()

	conn, err := net.Dial("tcp", addr.String())
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write([]byte("TESTING 1 2 3"))
	if err != nil {
		t.Error(err)
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}

	l.Close()

	<-done
}
