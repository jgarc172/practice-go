package main

import (
	"fmt"
	"io"
	"os"
)

/*
 program reads from stdin
 reads 16 bytes at a time
 and prints to stdout
 redirection can be used
 use 'ctrl+d' for EOF
 examples:
  $ go run .
  $ go run . > in.txt
  $ go run . < in.txt > out.txt
*/
func main() {
	readBytes(os.Stdin)
}

func readBytes(r io.Reader) {
	b := make([]byte, 16)
	for {
		n, err := r.Read(b)
		if n > 0 {
			fmt.Print(string(b[:n]))
		}
		if err != nil || err == io.EOF {
			break
		}
	}
}
