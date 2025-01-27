package main

import "fmt"

func main() {
	//Practice using for cycle
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	name := "Eduardo"
	for i := 0; i < len(name); i++ {
		fmt.Println(i, string(name[i]))
	}
	//WHILE does not exists in GOLAND
	userAge := 5
	for userAge < 18 {
		fmt.Println("Since you are ", userAge, " years old, you can´t access to this website.")
		userAge++
	}
	if userAge >= 18 {
		fmt.Println("Welcome to the Go Programming language!")
	}
}
