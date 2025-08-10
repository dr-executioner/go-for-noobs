package typesystem

import "fmt"


func Describe(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	case bool:
		fmt.Println("bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}
