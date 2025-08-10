package basics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Structs() {
	//Like a TypeScript type or interface with actual data
	person := Person{"Aarjya", 25}
	fmt.Println(person)

	fmt.Println(person.Name)
}
