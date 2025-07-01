package basics

import "fmt"

func Variables() {
	var name string = "Ronaldo"
	fmt.Println("Name:", name)

	// type inference
	age := 25
	fmt.Println("Age:", age)

	//multiple declerarions
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	//Constants
	const PI = 3.14
	fmt.Println("Pi:", PI)
}
