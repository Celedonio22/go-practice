package main

import "fmt"

// Declare an object of Person type
type person struct {
	name     string
	lastname string
	age      int
}

func main() {
	person1 := person{"Eduardo", "Garcia", 47}
	person2 := person{"Eduardo", "Celedonio", 28}
	fmt.Println(person1)
	fmt.Println(person2)
	person2.age = 29
	fmt.Println(person2)
}
