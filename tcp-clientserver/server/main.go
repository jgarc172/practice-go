package main

import (
	"fmt"
	"io"
	"log"
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
		session, err := newSession("bash")
		if err != nil {
			fmt.Println("could not start shell", "bash")
			conn.Close()
			continue
		}
		go attachTerminal(conn, session)
	}
}

// newSession starts a shell and returns the pseudo Terminal device
func newSession(shell string) (s session, err error) {
	cmd := exec.Command(shell)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return
	}
	s = session{ptmx, cmd.Process.Pid}
	return
}

// attachTerminal connects conn and pseudo terminal
// until connection is closed
func attachTerminal(conn net.Conn, s session) {
	defer func() {
		conn.Close()
		s.Close()
	}()
	log.Println("shell session started, pid:", s.pid)
	fmt.Fprint(conn, "welcome", "\r\n")

	go func() {
		io.Copy(s.pty, conn)
	}()
	io.Copy(conn, s.pty)

	fmt.Fprint(conn, "bye\r\n")
	log.Println("shell session ended, pid:", s.pid)
}

type session struct {
	pty *os.File
	pid int
}

func (s session) Close() {
	s.pty.Close()
}
