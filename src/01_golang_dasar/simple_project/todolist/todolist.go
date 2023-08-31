package main

import "fmt"

var todos = []map[string]string{
	{"judul": "Belajar Golang", "isi": "Belajar golang dasar sampai mahir"},
}

func tambah(judul, isi string) {
	todo := map[string]string{}
	todo["judul"] = judul
	todo["isi"] = isi
	todos = append(todos, todo)
}

func lihat() {
	for _, todo := range todos {
		fmt.Println("Judul: ", todo["judul"])
		fmt.Println("Isi: ", todo["isi"])
		fmt.Println("")
	}
}

func main() {
	var judul, isi string
	var pilihan int
	for {
		// Menu
		println("1. Tambah Todo")
		println("2. Lihat Todo")
		println("3. Keluar")

		// Input pilihan
		println("Pilihan: ")
		fmt.Scanln(&pilihan)

		// Cek pilihan
		switch pilihan {
		case 1:
			println("Judul: ")
			fmt.Scanln(&judul)
			println("Isi: ")
			fmt.Scanln(&isi)
			tambah(judul, isi)
		case 2:
			lihat()
		case 3:
			return
		default:
			println("Pilihan tidak tersedia")
		}
	}

}
