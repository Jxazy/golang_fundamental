package main

import "fmt"

func main() {

	// Constant adalah variable yang nilainya tidak bisa diubah
	const (
		name = "Fajar Octhaviano"
		age  = 20
		suka = true
	)

	fmt.Println("Nama saya", name)
	fmt.Println("Umur saya", age, "tahun")
	if suka == true {
		fmt.Println("Saya suka belajar Golang")
	} else {
		fmt.Println("Saya tidak suka belajar Golang")
	}

}
