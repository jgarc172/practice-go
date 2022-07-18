package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLines := ParseArgs(os.Args)
	fmt.Println(count(os.Stdin, countLines))
}

func count(r io.Reader, countLines bool) int {
	cnt := 0
	sc := bufio.NewScanner(r)
	if !countLines {
		sc.Split(bufio.ScanWords)
	}
	for sc.Scan() {
		cnt++
	}
	return cnt
}

func ParseArgs(args []string) (lines bool) {
	if len(args) < 2 {
		return
	}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.BoolVar(&lines, "l", false, "count lines")
	flags.Parse(args[1:])
	return
}
