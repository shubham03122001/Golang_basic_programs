package main

import "fmt"

func main() {
	Number := 145
	Value := Number

	sum := 0
	for Number != 0 {
		factorial := 1
		Remainder := Number % 10
		for i := 1; i <= Remainder; i++ {
			factorial *= i

		}
		sum += factorial
		Number /= 10
	}

	if Value == sum {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
