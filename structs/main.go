package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Houston",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//& gives you the memory address of the value this variable is pointing at
	// jimPointer := &jim
	jim.updateName("Dave")
	jim.print()
}

// In the function use it just means you are working with a point inside the func
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// The * operator in line 35 gives you the value this memory address is point at

func (p person) print() {
	fmt.Printf("%+v", p)
}
