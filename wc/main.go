package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
$ go build; ls wc
wc
Examples:
$ ./wc
hello world
<Ctrl+D>
2
$ echo "hello world" | ./wc
2
$ ./wc < main.go
61
*/
func main() {
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	wc := 0
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		wc++
	}
	return wc
}
