package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

/*
 go routine --> |---------------------------------------------|
                | pty-master----Line-Dispcipline-----pty-slave| <-->  cmd
     Stdout <-- | --------------------------------------------|
*/
func main() {
	c := exec.Command("grep", "--color=auto", "bar")
	// f is the pty-master
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte("foo\n"))
		f.Write([]byte("bar in this line\n"))
		f.Write([]byte("baz\n"))
		f.Write([]byte{4}) // EOT
	}()
	io.Copy(os.Stdout, f)
}
