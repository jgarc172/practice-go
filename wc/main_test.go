package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func ExampleCount() {
	reader := strings.NewReader("Hello World")
	fmt.Println(count(reader))
	// Output:
	// 2
}

// tests count(r io.Reader) int
func TestCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := count(tc.input)

			if got != tc.want {
				t.Errorf("got '%d' want '%d'", got, tc.want)
			}
		})
	}
}

var testCases = []struct {
	name  string
	input io.Reader
	want  int
}{
	{"empty", strings.NewReader(""), 0},
	{"one line", strings.NewReader("one two three\n"), 3},
	{"lines", strings.NewReader("one two three\nfour five\n"), 5},
}
