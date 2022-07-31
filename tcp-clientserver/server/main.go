package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

// go build -o server
// server <port #>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: server <port #>")
		os.Exit(1)
	}
	server, err := startServer(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("server started . . .")

	fmt.Println("use 'Ctrl+c' to stop it")
	serveShellConnections(server)

	fmt.Println("server exited")
}

func startServer(port string) (l net.Listener, err error) {
	addr := fmt.Sprintf(":%v", port)
	fmt.Println("will start server on address:port", addr)
	l, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return l, err
}

func serveShellConnections(server net.Listener) {
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}
		pty, err := startShell("bash")
		if err != nil {
			fmt.Println("could not start shell", "bash")
			conn.Close()
			continue
		}
		go attachTerminal(conn, pty)
	}
}

// startShell starts a shell and returns the pseudo Terminal device
func startShell(cmd string) (ptmx *os.File, err error) {
	shell := exec.Command(cmd)
	ptmx, err = pty.Start(shell)
	if err != nil {
		return nil, err
	}
	fmt.Println("shell session started with pid:", shell.Process.Pid)
	return ptmx, nil
}

// attachTerminal connects conn and pseudo terminal
// until connection is closed
func attachTerminal(conn net.Conn, pty *os.File) {
	defer func() {
		conn.Close()
		pty.Close()
	}()
	fmt.Fprint(conn, "welcome", "\r\n")

	go func() {
		io.Copy(pty, conn)
	}()
	io.Copy(conn, pty)

	fmt.Fprint(conn, "bye\r\n")
	fmt.Println("shell session ended")
}
