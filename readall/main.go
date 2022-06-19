package main

import (
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
	readAll(os.Stdin, os.Stdout)
}

func readAll(r io.Reader, w io.Writer) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(bytes))
}
