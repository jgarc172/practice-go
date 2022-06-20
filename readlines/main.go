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
	readLines(os.Stdin, os.Stdout)
}

// reads input line by line
// and prints it to output
func readLines(r io.Reader, w io.Writer) {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Fprint(w, line)
	}
}
