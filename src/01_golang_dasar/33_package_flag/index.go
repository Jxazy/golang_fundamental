package main

import (
	"flag"
	"fmt"
)

func main() {

	var hostname *string = flag.String("hostname", "localhost", "put your hostname")
	var username *string = flag.String("username", "root", "put your username")
	var password *string = flag.String("password", "root", "put your password")

	flag.Parse()

	fmt.Println("hostname: ", *hostname)
	fmt.Println("username: ", *username)
	fmt.Println("password: ", *password)

	fmt.Println("============")
	fmt.Println("Sisanya explore di https://pkg.go.dev/flag")

}
