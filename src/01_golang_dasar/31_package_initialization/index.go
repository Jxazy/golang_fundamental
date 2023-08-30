package main

import (
	"fmt"
	helper "golang_dasar/29_helper"
)

func main() {

	koneksiSaya := helper.MyConnection()
	fmt.Println(koneksiSaya)
}

// apabila ingin memanggil function init saja bisa menggunakan blank identifier _
/* contoh :

import (
	_ "golang_dasar/29_helper"
)

*/
