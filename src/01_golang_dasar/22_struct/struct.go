package main

import "fmt"

type User struct {
	Name, Address string
	Age           int
}

// Struct Method atau Function
func (fajar User) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", fajar.Name)
}

func main() {

	// Struct biasa
	fajar := User{
		Name:    "Fajar",
		Address: "Indonesia",
		Age:     22,
	}

	fmt.Println(fajar.Name)
	fmt.Println(fajar.Address)
	fmt.Println(fajar.Age)
	fmt.Println("=============")
	// memanggil struct method atau function
	fajar.sayHi("Dava")

}
