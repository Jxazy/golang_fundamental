package main

import "fmt"

func main() {
	fmt.Println("====== BREAK ======")
	for i := 0; i < 10; i++ {

		fmt.Println("Perulangan ke", i)
		if i == 5 {
			break
		}
	}

	fmt.Println("====== CONTINUE ======")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("====== BREAK NESTED LOOP ======")
	outerLoop := 0
	for {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break
			}
			fmt.Println("Perulangan ke", j)
		}
		if outerLoop == 2 {
			break
		}
		outerLoop++
	}

	fmt.Println("====== BREAK LABEL ======")
outerLoop2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop2
			}
			fmt.Println("Perulangan ke", i)
		}
	}

	fmt.Println("====== CONTINUE LABEL ======")
outerLoop3:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				continue outerLoop3
			}
			fmt.Println("Perulangan ke", i)
		}
	}

}
