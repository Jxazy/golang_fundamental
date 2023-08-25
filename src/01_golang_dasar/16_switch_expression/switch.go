package main

import "fmt"

func main() {

	name := "Fajar"

	switch name {
	case "Fajar":
		fmt.Println("Fajar Octhaviano")
	case "Budi":
		fmt.Println("Budi Santoso")
	default:
		fmt.Println("Siapa namanya?")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length >= 5:
		fmt.Println("Nama  terdiri dari 5 huruf atau lebih")
	default:
		fmt.Println("Nama sudah benar")
	}

}