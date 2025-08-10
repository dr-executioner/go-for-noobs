package typesystem

import "fmt"

func TypeAssert() {
	var i interface{} = "Aarjya"

	value, ok := i.(string)
	if ok {
		fmt.Println("Value is a string:", value)
	} else {
		fmt.Println("Not a string")
	}
}

// Panic Assertion
///📖 Panic on Failed Assertion (if no ok check)
/// If you omit ok and it fails, Go will panic.
//var i interface{} = 42
//value := i.(string) // runtime panic!
