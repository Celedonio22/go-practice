package main

import "fmt"

// Declare an object of Person type
type person struct {
	name     string
	lastname string
	age      int
}

func (p person) greet(greeting string) string {
	return greeting + ", " + p.name
}
func (p person) incrementAge() int {
	p.age++
	return p.age
}

func main() {
	person1 := person{"Eduardo", "Garcia", 47}
	person2 := person{"Eduardo", "Celedonio", 28}
	fmt.Println(person1)
	fmt.Println(person2)
	person2.age = 29
	fmt.Println(person2)
	greeting := person1.greet("Hola")
	fmt.Println(string(greeting))
	fmt.Println(person1.incrementAge())
}
