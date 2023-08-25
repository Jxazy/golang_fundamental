package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Fajar",
		"address": "Indonesia",
		"job":     "Programmer",

	}

	// menambah data ke map
	person["age"] = "20"

	fmt.Println("Nama Saya", person["name"])
	fmt.Println("Saya tinggal di", person["address"])
	fmt.Println("saat ini saja bekerja sebagai", person["job"])
	fmt.Println("Umur saya", person["age"], "Tahun")
	fmt.Println("=====================================")

	// membuat data map kosong
	var dataMahasiswa map[string]string = make(map[string]string)

	dataMahasiswa["nama"] = "Fajar"
	dataMahasiswa["nim"] = "123456789"
	dataMahasiswa["fakultas"] = "FST"
	dataMahasiswa["jurusan"] = "Sistem Informasi"

	// menghapus data map
	delete(dataMahasiswa, "fakultas")

	fmt.Println(dataMahasiswa)

}