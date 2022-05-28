package main

import (
	"bufio"
	"io"
)

func main() {

}

func count(r io.Reader) int {
	wc := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wc++
	}
	return wc
}
