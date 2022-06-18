package main

import (
	"fmt"
	"os"
)

// programm enforces use of two arguments
// ex: go run . one two
func main() {
	if len(os.Args) != 3 {
		fmt.Println("error: two arguments not provided")
		return
	}
	fmt.Printf("first argument: %s\n", os.Args[1])
	fmt.Printf("second argument: %s\n", os.Args[2])
}
