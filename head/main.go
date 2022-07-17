package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	usage := "usage: ./head address:port"
	address := arg(os.Args)
	if address == "" {
		fmt.Println(usage)
		return
	}
	h, err := Head(address)
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println(usage)
		return
	}
	fmt.Println(h)
}

func arg(args []string) (first string) {
	if len(args) > 1 {
		first = args[1]
	}
	return
}
func Head(addr string) (head string, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		return
	}
	resp, err := ioutil.ReadAll(conn)
	if err != nil {
		return
	}
	return string(resp), nil
}
