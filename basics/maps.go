package basics

import "fmt"

func Maps() {

	// Create a map
	person := map[string]string{
		"name": "Aarjya",
		"city": "Rayagada",
	}

	// access value
	fmt.Println(person["name"])

	// add key-value pair
	person["age"] = "25"

	//loop over map
	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
