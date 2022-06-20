package main

import (
	"io"
	"os"
)

/* copy in (stdin) to out (stdout)
examples with redirection:
 no redirection. copies what you type and is echoed
 $ go run .
 copies what you type to in.txt file
 $ go run . > in.txt
 $ cat in.txt
 copies in.txt to out.txt
 $ go run . < in.txt > out.txt
 $ cat out.txt
*/
func main() {
	copy(os.Stdout, os.Stdin)
}

// copies contents from Reader to Writer
func copy(w io.Writer, r io.Reader) {
	io.Copy(w, r)
}
