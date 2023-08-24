package main

import "fmt"

func main() {

	// Operasi Matematika
	var (
		penjumlahan = 10 + 10
		pengurangan = 10 - 5
		perkalian   = 10 * 5
		pembagian   = 10 / 5
		modulo      = 10 % 5
	)

	fmt.Println("Hasil penjumlahan:", penjumlahan)
	fmt.Println("Hasil pengurangan:", pengurangan)
	fmt.Println("Hasil perkalian:", perkalian)
	fmt.Println("Hasil pembagian:", pembagian)
	fmt.Println("Hasil modulo:", modulo)
	fmt.Println("====================")

	// Augmented Assignment
	var nilai = 10
	nilai += 10 // nilai = nilai + 10

	fmt.Println("Nilai =", nilai)
	fmt.Println("====================")

	// Unary Operator
	nilai++ // nilai = nilai + 1
	fmt.Println("Nilai =", nilai)

}
