package main

import "fmt"

func main() {
	// for loop basic
	perulangan := 1
	for perulangan <= 5 {
		fmt.Println("Perulangan ke", perulangan)
		perulangan++
	}

	fmt.Println("=====For Loop 2 Statement=========")
	// for loop dengan 2 statement
	// for init statement; condition; post statement {
	for statement := 1; statement <= 5; statement++ {
		fmt.Println("Statement ke", statement)
	}

	fmt.Println("=======For Range Slice==========")
	// for range dengan slice
	var kelas = [...]string{
		"X MIPA",
		"X IPS",
		"XI MIPA",
		"XI IPS",
		"XII MIPA",
		"XII IPS",
	}

	for index, kelas := range kelas {
		fmt.Println("Index ke", index, "=", kelas)
	}
	fmt.Println("=====For Range tanpa Index==========")
	// Apabila tidak memerlukan index
	for _, kelas := range kelas {
		fmt.Println(kelas)
	}

	fmt.Println("=====For Range Map==========")
	// for range dengan map
	var kelasMap = map[string]string{
		"Fajar": "X MIPA",
		"Aldi":  "X IPS",
		"Rizky": "XI MIPA",
		"Raka":  "XI IPS",
	}

	for key, value := range kelasMap {
		fmt.Println(key, "Kelas", value)
	}

	fmt.Println("=====For Range Map tanpa Key==========")
	// Apabila tidak memerlukan key
	for _, value := range kelasMap {
		fmt.Println(value)
	}

}
