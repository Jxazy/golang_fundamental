package main

import "fmt"

func main() {

	var angka1, angka2 float64
	var operator string
	fmt.Println("Enter first number: ")
	fmt.Scanln(&angka1)
	fmt.Println("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&angka2)
	var result float64
	switch operator {
	case "+":
		result = angka1 + angka2
	case "-":
		result = angka1 - angka2
	case "*":
		result = angka1 * angka2
	case "/":
		result = angka1 / angka2
	default:
		fmt.Println("Invalid operator")
		return
	}
	fmt.Printf("%.2f %s %.2f = %.2f\n", angka1, operator, angka2, result)
}
