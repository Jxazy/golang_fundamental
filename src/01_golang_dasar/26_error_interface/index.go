package main

import (
	"errors" // Menggunakan package errors
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil pembagian:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
