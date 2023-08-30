package main

import (
	"fmt"
	helper "golang_dasar/29_helper"
)

func main() {
	helper.SayHello("Fajar") // Dapat mengimport dari package lain karena diawali dengan huruf besar
	// helper.sayGoodbye tidak bisa karena diawali dengan huruf kecil (tidak terbaca)
	fmt.Println(helper.NamaLengkap)
}
