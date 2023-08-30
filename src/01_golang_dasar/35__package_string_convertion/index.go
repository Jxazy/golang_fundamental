package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Konversi dari String ke boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Konversi dari String ke Int
	string, err := strconv.ParseInt("22", 10, 32)
	if err == nil {
		fmt.Println(string)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Konversi dari String ke Float
	float, err := strconv.ParseFloat("22.5", 64)
	if err == nil {
		fmt.Println(float)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Cara cepat menggunakan Atoi (String menjadi Integer)
	valueInt, err := strconv.Atoi("2000000")
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Cara cepat menggunakan Itoa (Int menjadi String)
	valueString := strconv.Itoa(4000000)
	fmt.Println(valueString)

	fmt.Println("============")
	fmt.Println("Sisanya explore di https://pkg.go.dev/strconv")
}
