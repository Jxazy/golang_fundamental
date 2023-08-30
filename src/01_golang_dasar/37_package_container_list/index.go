package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	data.PushBack("Fajar")
	data.PushBack("Octhaviano")
	data.PushFront("Muhammad")

	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	// full doc: https://pkg.go.dev/container/lists

}
