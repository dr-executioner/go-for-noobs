package basics

//function with return type int
func Add(a int, b int) int {
	return a + b
}


// function returning multiple values
func Divide(a, b int) (int, string) {
	if b == 0 {
		return 0, "Cannot divide by zero"
	}
	return a / b, "Success"
}
