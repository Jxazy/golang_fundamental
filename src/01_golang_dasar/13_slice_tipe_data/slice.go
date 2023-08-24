package main

import "fmt"

func main() {

	const pembatas = "===================="

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"Mei",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"Desember",
	}

	var slice1 = months[10:]
	fmt.Println(slice1)
	fmt.Println("length:", len(slice1))
	fmt.Println("capacity:", cap(slice1))
	fmt.Println(pembatas)

	// append = mereplace data yang ada di slice, apabila data yang di append lebih dari capacity maka akan membuat array baru
	var slice2 = append(slice1, "fajar")
	fmt.Println(slice2)
	fmt.Println(months)
	fmt.Println(pembatas)

	// make = membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fajar"
	newSlice[1] = "Octhaviano"
	fmt.Println(newSlice)
	fmt.Println(pembatas)

	// Copy slice source ke slice destination
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)
	fmt.Println(pembatas)

}
