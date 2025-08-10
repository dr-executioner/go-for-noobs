package basics

import "fmt"

func Maker() {

	// make a slice
	numbers := make([]int, 3)
	numbers[0] = 10
	fmt.Println(numbers)
	
	//make a map
	person := make(map[string]string)
	person["age"]= "23"
	fmt.Println(person)
}