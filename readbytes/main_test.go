package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// TestReadBytes tests
// readBytes(r io.Reader, w io.Writer)
func TestReadBytes(t *testing.T) {
	want := "line 1\nline 2\n"

	r := newReader(want)
	w := newWriter()
	readBytes(r, w)
	got := w.String()

	if got != want {
		t.Errorf("Not Empty: got '%s' want '%s'", got, want)
	}

	want = ""
	r = newReader(want)
	w = newWriter()
	readBytes(r, w)
	got = w.String()

	if got != want {
		t.Errorf("Empty: got '%s' want '%s'", got, want)
	}

}

func newReader(s string) io.Reader {
	return strings.NewReader(s)
}

func newWriter() *bytes.Buffer {
	return &bytes.Buffer{}
}
