package main

import (
	"context"
	"fmt"
	"myapp/advanced"
	"time"
)

func main() {
	// result, err := basics.DivideCheck(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Division result:", result)

	// // Error Patterns Usage
	// item, err := errorpatterns.GetItem(6)
	// if err != nil {
	// 	if errors.Is(err, errorpatterns.ErrNotFound) {
	// 		fmt.Println("Handled Sentinel Error:", err)
	// 		return
	// 	}
	// 	fmt.Println("Unhandled Error:", err)
	// 	return
	// }
	// fmt.Println("Found Item:", item)

	// //Custom Error Usage
	// vAge := errorpatterns.ValidateAge(90)
	// if vAge != nil {
	// 	if vErr, ok := vAge.(errorpatterns.ValidationError); ok {
	// 		fmt.Println("Validation Error", vErr.Msg)
	// 	} else {
	// 		fmt.Println("Other error")
	// 	}
	// 	return
	// }
	// fmt.Println("Age is", vAge)

	// // Error Wrapper Usage
	// fetchDataErr := errorpatterns.FetchData()
	// if fetchDataErr != nil {
	// 	if errors.Is(fetchDataErr, errorpatterns.ErrDatabase) {
	// 		fmt.Println("Caught database error inside wrapped error")
	// 	} else {
	// 		fmt.Println("Other Error:", err)
	// 	}
	// }

	// // Type Switch uses
	// typesystem.Describe("Aarjya")
	// typesystem.Describe(69)
	// typesystem.Describe(true)

	// Goroutines
	// ch:= make(chan string)

	// go func(){
	// 	ch <- "Channel is live"
	// }()

	// msg,ok := <-ch
	// fmt.Println("Channel Status", ok)
	// go concurrencyfundamentals.Greet("Arya", msg)

	// //closing channel
	// close(ch)
	// fmt.Println("Channel Status", ok)

	// fmt.Println("Hello, i am Main")
	// time.Sleep(1*time.Second)

	// Goroutines and Channels example
	// concurrencyfundamentals.Factory()

	//Context
	// Step 1: create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())

	// Step 2: start worker
	for i := 1; i < 5; i++ {
		go advanced.Worker(ctx, i)
		go advanced.Worker(ctx2, i)
		go advanced.Worker(ctx, i)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Main: stopping ONLY worker 2 (nested context cancel)")
	cancel2()
	
	// Step 3: let it run for some time
	time.Sleep(5 * time.Second)

	// Step 4: cancel the context
	fmt.Println("Main: canceling context")
	cancel()

	// Give goroutine time to exit
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main: done")
}
