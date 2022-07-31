package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: client <address:port>")
		os.Exit(1)
	}
	client, err := startClient(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()

	attachTerminal(client)
}

func attachTerminal(client net.Conn) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("failed to make 'raw', this client terminal")
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	go func() {
		io.Copy(client, os.Stdin)
	}()
	io.Copy(os.Stdout, client)
}

func startClient(addr string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", addr)
	return
}
