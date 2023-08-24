package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var ktpFajar noKTP = "1234567890"
	var marriedStatus married = false

	fmt.Println("Nomor KTP fajar adalah:", ktpFajar)
	if marriedStatus == true {
		fmt.Println("Fajar sudah menikah")
	} else {
		fmt.Println("Fajar belum menikah")
	}

}
