package main

import "fmt"

func main() {
	fmt.Println(add(10, 5))
	fmt.Println(substraction(10, 5))
	fmt.Println(division(10, 5))
	fmt.Println(greetings("Eduardo"))
}

func add(x int, y int) int {
	return x + y
}
func substraction(x int, y int) int {
	return x - y
}
func division(x int, y int) int {
	return x / y
}
func greetings(name string) string {
	return "Hello, " + name + "!"
}
