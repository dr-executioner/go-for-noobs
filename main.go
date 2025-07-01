package main

import (
	"fmt"
	"myapp/basics"
)

func main() {
	result, err := basics.DivideCheck(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Division result:", result)
}
