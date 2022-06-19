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
	readLines(os.Stdin)
}

// reads input line by line
// and prints it to stdout
func readLines(r io.Reader) {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}
