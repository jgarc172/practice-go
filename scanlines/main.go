package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
 program to copy stdin to stdout with redirection
 examples:
  $ go run .
  $ go run . > in.txt
  $ go run . < in.txt > out.txt
*/
func main() {
	scanLines(os.Stdin, os.Stdout)
}

func scanLines(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Fprintln(w, sc.Text())
	}
}
