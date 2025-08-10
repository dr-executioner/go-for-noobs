package typesystem

import "fmt"

type Celsius float64

//If you define a new type, you can attach methods to it.
func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func CustomTypes() {
	var temp Celsius = 30
	fmt.Println(temp.ToFahrenheit())
}
