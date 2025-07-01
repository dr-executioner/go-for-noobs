package basics

import "fmt"

func DeferPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("Recovered from panic:", r)
		}
	}()

	defer fmt.Println("Deferred message 1")
	defer fmt.Println("Deferred message 2")

	fmt.Println("Before Panic")
	panic("Something went seriously wrong")
	fmt.Println("After Panic") // this line is never reached

}
