package main

import "fmt"

func main() {
	adition := 5 + 10
	fmt.Println(adition)
	substraction := 10 - 5
	fmt.Println(substraction)
	multiplication := 5 * 5
	fmt.Println(multiplication)
	division := 10 / 5
	fmt.Println(division)
	modulus := 10 % 5
	fmt.Println(modulus)
	//increment
	adition++
	fmt.Println(adition)
	//decrement
	adition--
	fmt.Println(adition)

	//comparision operators
	fmt.Println(10 < 5)
	fmt.Println(10 > 5)
	fmt.Println(10 <= 5)
	fmt.Println(10 >= 5)
	fmt.Println(10 == 5)
	fmt.Println(10 != 5)

	fmt.Println(10 != 5 && 10 != 10)
	fmt.Println(10 != 5 || 10 == 10)

}
