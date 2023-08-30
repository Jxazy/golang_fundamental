package main

import "fmt"

func asal() interface{} {
	return false
}

func main() {

	var merubahData interface{} = asal()

	switch value := merubahData.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Int")
	default:
		fmt.Println("Unknown Value:", value)

	}

}

// Type Assertion berfungsi untuk merubah data dari interface menjadi tipe data lain
