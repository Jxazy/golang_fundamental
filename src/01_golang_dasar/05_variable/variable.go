package main

import "fmt"

func main() {
	var pembatas = "===================="

	// Variable Dasar
	var name string

	name = "Fajar Octhaviano"
	fmt.Println("Nama saya", name)
	fmt.Println(pembatas)

	name = "Fajar.ahmdd"
	fmt.Println("Usernbame instagram saya", name)
	fmt.Println(pembatas)

	// Variable Default value
	var age int8 = 20

	fmt.Println("Umur saya", age, "tahun")
	fmt.Println(pembatas)

	age = 21
	fmt.Println("Umur saya tahun depan", age, "tahun")
	fmt.Println(pembatas)

	// Deklarasi Variable tanpa var
	profesi := "Back End Developer"
	fmt.Println("Saya seorang", profesi)
	fmt.Println(pembatas)

	// Deklarasi Variable tanpa var dengan multi variable
	var (
		firstName = "Fajar"
		lastName  = "Octhaviano"
	)

	fmt.Println("Nama saya", firstName, lastName)

}
