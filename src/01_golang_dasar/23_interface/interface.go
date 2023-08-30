package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Contract #1
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Contract #2
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var adit Person
	adit.Name = "Adit"
	SayHello(adit)

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}
