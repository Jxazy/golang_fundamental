package main

import "fmt"

func main() {
	name := "Fajar Octhaviano"

	dataDiri := func() {
		// Closure akan mengubah variabel yang berada di luar scope-nya
		// Agar tidak terjadi perubahan variabel, maka harus di deklarasikan ulang menggunakan :=
		name = "Fajar Ahmad"

		fmt.Println("Nama saya", name)

	}

	dataDiri()
	fmt.Println("Nama saya", name)

}
