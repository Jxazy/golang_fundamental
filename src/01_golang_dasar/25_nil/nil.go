package main

import "fmt"

// nil hanya support di interface, function, map, slice, pointer dan channel.

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("Fajar")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

}
