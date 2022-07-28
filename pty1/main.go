package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

/*
 go routine -->Write() |In----------------------------------------Out| Read()
                       | pty-master----Line-Discipline-----pty-slave | <------> cmd
     Stdout <--Read()  |Out----------------------------------------In| Write()

In  -> Writer, can write to, writable
Out -> Reader, can read from, readable
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
