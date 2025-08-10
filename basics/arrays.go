package basics

import "fmt"

func Arrays() {
	// fixed size collections of same type
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println(numbers)

	// Access Elements
	fmt.Println(numbers[0])

	//Length
	fmt.Println(len(numbers))

}

func Slices() {
	fruits := []string{"Apple", "Banana", "Cherry"}

	// Add New fruits
	fruits = append(fruits, "Mango")

	// Access
	fmt.Println(fruits[0])

	//Length
	fmt.Println(len(fruits))

	//Sub-Slicing
	sub:=fruits[1:3]
	fmt.Println(sub)
}
