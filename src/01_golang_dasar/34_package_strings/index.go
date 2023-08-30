package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fajar Octhaviano", "Fajar"))
	fmt.Println(strings.Split("Fajar Octhaviano", " "))
	fmt.Println(strings.ToLower("Fajar Octhaviano"))
	fmt.Println(strings.ToUpper("Fajar Octhaviano"))
	fmt.Println(strings.Trim("      Fajar Octhaviano        ", " "))
	fmt.Println(strings.ReplaceAll("Fajar Fajar Fajar ganteng", "Fajar", "Iqbal"))

	fmt.Println("============")
	fmt.Println("Sisanya explore di https://pkg.go.dev/strings")

}
