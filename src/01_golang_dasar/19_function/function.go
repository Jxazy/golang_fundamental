package main

import "fmt"

// function basic

func nama() {
	fmt.Println("Hallo, nama saya Fajar Octhaviano")
}

// function dengan parameter

func age(age int) {
	fmt.Println("Umur saya", age, "tahun")
}

// function dengan return value

func pendidikan(age int) string {
	if age >= 12 && age <= 15 {
		return "SMP"
	} else if age >= 16 && age <= 18 {
		return "SMA"
	} else if age >= 19 && age <= 22 {
		return "Kuliah"
	} else {
		return "Tidak tahu"
	}
}

// function dengan return multiple value

func tempatTinggal() (string, string) {
	return "Bekasi,", "Jawa Barat."
}

// function multiple hanya mengambil salah satu value & named return value

func detailAlamat() (kota string, durasi string) {
	kota = "Bekasi"
	durasi = "3 tahun"

	return // named return value tidak wajib ditulis
}

// function dengan parameter variadic

func hitungAngka(tahun ...int) int {

	tahunMasuk := tahun[0]
	tahunKeluar := tahun[0]

	for _, num := range tahun {
		if num < tahunMasuk {
			tahunMasuk = num
		}
		if num > tahunKeluar {
			tahunKeluar = num
		}
	}
	return tahunKeluar - tahunMasuk
}

// Function value & as parameter

type Filter func(string) string

func lulus(status string, filter Filter) {
	statusLulus := filter(status)
	fmt.Println("Status:", statusLulus)
}

func statusSMA(status string) string {
	if status == "Lulus" {
		return "telah lulus SMA"
	} else {
		return "belum lulus SMA "
	}
}

// Anonymous Function

type Status func(int) bool

func registerUser(semester int, status Status) {
	if status(semester) {
		fmt.Println("Kamu belum lulus, kamu berada di semester", semester)
	} else {
		fmt.Println("Selamat Kamu Sudah lulus")
	}
}

// Recursive Function

func faktorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	// function basic
	nama()
	// function dengan parameter
	age(20)
	// function dengan return value
	fmt.Println("Saya berada di jenjang", pendidikan(20))
	// function dengan return multiple value
	kota, provinsi := tempatTinggal()
	fmt.Println("Saya tinggal di", kota, "Provinsi", provinsi)
	// function multiple hanya mengambil salah satu value & named return value
	_, durasi := detailAlamat()
	fmt.Println("Saya tinggal di sana selama", durasi)
	// variadic parameter
	selisih := hitungAngka(2020, 2023)
	fmt.Println("Total saya belajar di kampus adalah selama", selisih, "tahun")
	// Function value
	lulus("Lulus", statusSMA)
	// Anonymous Function
	status := func(semester int) bool {
		return semester >= 1 && semester <= 8
	}
	registerUser(7, status)
	// Recursive Function
	fmt.Println("Hasil dari 5! adalah", faktorial(5))

}
