package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	// Jika nilai8 > 127, maka nilai8 akan kembali ke -128
	// Begitu juga sebaliknya, jika nilai8 < -128, maka nilai8 akan kembali ke 127
	// Hal ini disebut dengan overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println("=====================================")

	var name = "Fajar Octhaviano"
	var e = name[0]

	// Konversi dari byte ke string
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
