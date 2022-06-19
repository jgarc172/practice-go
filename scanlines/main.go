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
	scanLines(os.Stdin)
}

func scanLines(r io.Reader) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
