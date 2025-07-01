package basics

import "fmt"

func Conditionald(age int) {
	if age < 18 {
		fmt.Println("Minor")
	}else if age  == 18 {
		fmt.Println("Just became 18")
	}else{
		fmt.Println("Adult")
	}

	// switch
	switch {
	case age == 18 :
		fmt.Println("Eighteen")
	case age > 18:
		fmt.Println("Adult")
	default :
	fmt.Println("Minor")	
	}

}