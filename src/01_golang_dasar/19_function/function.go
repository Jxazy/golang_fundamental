package main

import "fmt"

// function basic
func nama() {
	fmt.Println("Hallo, nama saya Fajar Octhaviano")
}

// function dengan parameter
func age(age int) {
	fmt.Println("Umur saya", age, "tahun")
}

// function dengan return value
func pendidikan(age int) string {
	if age >= 12 && age <= 15 {
		return "SMP"
	} else if age >= 16 && age <= 18 {
		return "SMA"
	} else if age >= 19 && age <= 22 {
		return "Kuliah"
	} else {
		return "Tidak tahu"
	}
}

// function dengan return multiple value

func main() {
	nama()
	age(20)
	fmt.Println("Saya berada di jenjang", pendidikan(20))
}
