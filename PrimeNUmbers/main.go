package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the element")
	fmt.Scan(&number)
	count := 0
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 0 {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not prime")
	}
}
