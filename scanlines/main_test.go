package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// TestScanLines tests
// scanLines(r io.Reader, w io.Writer)
func TestScanLines(t *testing.T) {
	for _, test := range testData() {
		t.Run(test.name, func(t *testing.T) {

			scanLines(test.input1, test.input2)

			got := test.input2.String()
			if got != test.want {
				t.Errorf("got '%s' want '%s'", got, test.want)
			}
		})
	}
}

type test struct {
	name   string
	want   string
	input1 io.Reader
	input2 *bytes.Buffer
}

func testData() []test {
	want := []string{
		"",
		"one two three\n",
		"one two three\nfour five\n",
	}
	return []test{
		{"Empty", want[0], newReader(want[0]), newWriter()},
		{"Line", want[1], newReader(want[1]), newWriter()},
		{"Lines", want[2], newReader(want[2]), newWriter()},
	}
}

func newReader(s string) io.Reader {
	return strings.NewReader(s)
}

func newWriter() *bytes.Buffer {
	return &bytes.Buffer{}
}
