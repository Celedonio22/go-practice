package main

import "fmt"

/*
You need to iterate the numbers from 0 to 20 and you need to show only the even numbers
*/
func main() {
	findEvenNumbers(0, 20)
}
func findEvenNumbers(start int, end int) {
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println(i, " It´s a even number")
		}
	}
}
