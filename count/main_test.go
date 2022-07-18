package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func ExampleParseArgs() {
	args := []string{"count"}
	countLines := ParseArgs(args)
	fmt.Println(countLines)
	args = []string{"count", "-l"}
	countLines = ParseArgs(args)
	fmt.Println(countLines)
	// Output:
	// false
	// true
}

func ExampleCount() {
	reader := strings.NewReader("one two\nthree four")
	countLines := false
	fmt.Println(count(reader, countLines))
	reader = strings.NewReader("one two\nthree four")
	countLines = true
	fmt.Println(count(reader, countLines))
	// Output:
	// 4
	// 2
}
func TestParseArgs(t *testing.T) {
	for _, tc := range casesParseArgs {
		t.Run(tc.name, func(t *testing.T) {
			lines := ParseArgs(tc.args)
			if lines != tc.lines {
				t.Errorf("got '%v', wanted '%v'", lines, tc.lines)
			}
		})
	}
}

var casesParseArgs = []struct {
	name  string
	args  []string
	lines bool
}{
	{"default", []string{"./count"}, false},
	{"present", []string{"./count", "-l"}, true},
}

func TestCount(t *testing.T) {
	for _, tc := range casesCount {
		t.Run(tc.name, func(t *testing.T) {

			got := count(tc.reader, tc.lines)

			if got != tc.want {
				t.Errorf("got '%d' want '%d'", got, tc.want)
			}
		})
	}
}

var casesCount = []struct {
	name   string
	reader io.Reader
	lines  bool
	want   int
}{
	{"empty words", strings.NewReader(""), false, 0},
	{"three words", strings.NewReader("one two three\n"), false, 3},
	{"five words", strings.NewReader("one two three\nfour five\n"), false, 5},
	{"empty lines", strings.NewReader(""), true, 0},
	{"one line", strings.NewReader("one two three\n"), true, 1},
	{"two lines", strings.NewReader("one two three\nfour five\n"), true, 2},
}
