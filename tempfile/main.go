package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.CreateTemp(".", "temp*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())
}
