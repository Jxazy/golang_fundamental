package main

import "fmt"

func main() {

	var names [2]string

	names[0] = "Fajar"
	names[1] = "Octhaviano"

	fmt.Println(names[0], names[1])

	// array langsung
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	// array function
	var fruits = [3]string{"apple", "grape", "banana"}

	fmt.Println("Jumlah element \t\t", len(fruits))

}
