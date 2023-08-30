package main

type User struct {
	Name string
	age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func main() {

}
