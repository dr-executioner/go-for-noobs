package concurrencyfundamentals

import "fmt"

func Greet(name, msg string) {
	fmt.Println("Hello", name,".", msg)
}
