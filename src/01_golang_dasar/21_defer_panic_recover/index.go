package main

import "fmt"

func contohDefer() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("contohDefer = tetap dijalankan walaupun terjadi error")

}

func funcBasic(error bool) {
	defer contohDefer()
	fmt.Println("Aplikasi Dimulai")

	if error {
		panic("Fungsi Panic = memberikan informasi error dan menghentikan program")
	}
}

func main() {
	funcBasic(false)
}
