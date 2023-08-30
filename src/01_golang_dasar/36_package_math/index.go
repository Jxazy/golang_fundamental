package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
	fmt.Println(math.Pow(10, 2))
	fmt.Println(math.Cos(24))

	fmt.Println("============")
	fmt.Println("Sisanya explore di https://pkg.go.dev/math")
}
