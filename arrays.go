package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fruits := []string{"Apple", "Orange", "Pear"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
		if fruits[i] == "Apple" {
			fmt.Println("There is a coincidence!")
		}
	}
	fruits = append(fruits, "Banana", "Peach", "Strawberry")
	fmt.Println(fruits)
}
