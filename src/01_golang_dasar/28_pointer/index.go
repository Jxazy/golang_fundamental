package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Man struct {
	Name string
}

// Pointer Function
func changeAddressToIndonesia(address *Address) {
	address.Province = "Bali"
}

// Pointer Method (menggunakan pointer *)
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	var address1 Address = Address{"Bandar Lampung", "Lampung", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bekasi"

	// Untuk bisa merubah address 1 melalui address 2 maka harus diberikan * di awal syntax
	*address2 = Address{"Depok", "Jawa Barat", "Indonesia"}

	fmt.Println("=================")
	fmt.Println("Pointer di Variable:")
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address1.City)
	fmt.Println("=================")
	fmt.Println("Pointer NEW:")
	var address4 *Address = new(Address)
	fmt.Println(address4)
	fmt.Println("=================")
	fmt.Println("Pointer di Function:")

	// Pointer Function
	var alamat = Address{
		City:     "Denpasar",
		Province: "",
		Country:  "Indonesia",
	}

	changeAddressToIndonesia(&alamat)
	fmt.Println(alamat)
	fmt.Println("=================")
	fmt.Println("Pointer di Method:")

	// Pointer Method
	fajar := Man{"Fajar"}
	fajar.Married()
	fmt.Println(fajar.Name)

}
