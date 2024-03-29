package main

import (
	"fmt"
	"os"
)

// program gets and changes the process working directory
func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("process working directory: ", pwd)

	if err = os.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}
	pwd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final process working directory: ", pwd)
}
