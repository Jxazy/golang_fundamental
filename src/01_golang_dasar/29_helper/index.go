package helper

import "fmt"

// init adalah function yang pertama kali dijalankan secara otomatis
var connection string

func init() {
	fmt.Println("init berjalan tanpa dipanggil")
	connection = "MySQL"
}

// Dapat di import ke luar package
func SayHello(name string) {
	fmt.Println("Hello", name)
}

func MyConnection() string {
	return connection
}

var NamaLengkap = "Fajar Octhaviano"

// tidak dapat di import ke luar package
func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
