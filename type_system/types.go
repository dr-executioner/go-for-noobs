package typesystem

import "fmt"

type Meter float64            // New type
type Kilogram float64         // New type
type ID = int                 // Type alias

func TypeSystem() {
	var length Meter = 5.5
	var weight Kilogram = 70
	var userID ID = 101

	fmt.Println("Length:", length)
	fmt.Println("Weight:", weight)
	fmt.Println("UserID:", userID)

	// userID and int can be mixed because ID is an alias
	var num int = userID
	fmt.Println("Num:", num)

	// But Meter and Kilogram are distinct
	// var invalid Meter = weight  // this would give a compiler error
}
