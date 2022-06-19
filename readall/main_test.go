package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	want := "line 1\nline 2"

	r := newReader(want)
	w := newWriter()
	readAll(r, w)
	got := w.String()

	if got != want {
		t.Errorf("Not Empty: got '%s' want '%s'", got, want)
	}
}

func newReader(s string) io.Reader {
	return strings.NewReader(s)
}

func newWriter() *bytes.Buffer {
	return &bytes.Buffer{}
}
