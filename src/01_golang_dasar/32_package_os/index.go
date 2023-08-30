package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument: ")

	// Mengambil data argument
	fmt.Println(args[1])
	fmt.Println(args[2])

	fmt.Println("============")

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println("Hostname: ", err)
	}

	fmt.Println("============")
	fmt.Println("Sisanya explore di https://pkg.go.dev/os")
}
