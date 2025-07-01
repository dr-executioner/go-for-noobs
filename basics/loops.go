package basics

import "fmt"

func Loops(i int) {
	// for loop like while
	for i < 5 {
		fmt.Println("I is ", i)
		i++
	}

	// traditoonal for loop
	for i := 0; i < 6; i++ {
		fmt.Println("I is", i)
	}

	// for range loop (  over array )
	numbers := []int{10, 20, 30, 45, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value : %d\n", index, value)
	}
}
