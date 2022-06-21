package main

import (
	"io"
	"strings"
	"testing"
)

// tests count(r io.Reader) int
func TestCount(t *testing.T) {
	for _, test := range testData() {
		t.Run(test.name, func(t *testing.T) {
			got := count(test.input)
			if got != test.want {
				t.Errorf("got '%d' want '%d'", got, test.want)
			}
		})
	}
}

func newReader(s string) io.Reader {
	return strings.NewReader(s)
}

type test struct {
	name  string
	input io.Reader
	want  int
}

func testData() []test {
	return []test{
		{"empty", newReader(""), 0},
		{"one line", newReader("one two three\n"), 3},
		{"lines", newReader("one two three\nfour five\n"), 5},
	}
}
