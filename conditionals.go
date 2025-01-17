package main

import "fmt"

func main() {
	age := -8
	if age >= 18 {
		fmt.Println(age, "is greater than 18")
	} else if age < 18 && age >= 0 {
		fmt.Println(age, "is less than 18")
	} else {
		fmt.Println(age, "is not a valid age")
	}
}
