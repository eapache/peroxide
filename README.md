peroxide
========

[![GoDoc](https://godoc.org/github.com/eapache/peroxide?status.png)](https://godoc.org/github.com/eapache/peroxide)
[![Build Status](https://travis-ci.org/eapache/peroxide.svg?branch=master)](https://travis-ci.org/eapache/peroxide)

Peroxide provides some simple proxy-server-like functions in golang to simplify
testing reliable or redundant network services and bindings.

These projects have corner cases that are otherwise difficult to test as they
involve bringing the service up/down at specific intervals in order
to test the correct availability sequence. Peroxide makes this somewhat easier
to accomplish, as you don't have to touch the running service; simply
start/stop a proxy at will.
