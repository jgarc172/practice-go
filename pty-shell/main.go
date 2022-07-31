package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"golang.org/x/term"
)

/*
     Stdin  -->Write() |In----------------------------------------Out| Read()
                       | pty-master----Line-Discipline-----pty-slave | <------> shell
     Stdout <--Read()  |Out----------------------------------------In| Write()

In  -> Writer, can write to, writable
Out -> Reader, can read from, readable
*/
func main() {
	fmt.Println("starting bash . . .")
	defer func() {
		fmt.Println(". . . exited bash")
	}()
	c := exec.Command("bash")
	// ptmx is the pty-master
	ptmx, err := pty.Start(c)
	if err != nil {
		panic(err)
	}
	defer ptmx.Close()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() {
		term.Restore(int(os.Stdin.Fd()), oldState)
	}()

	go func() {
		io.Copy(ptmx, os.Stdin)
	}()
	io.Copy(os.Stdout, ptmx)
}
