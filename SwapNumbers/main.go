package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter the two numbers")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Println("Elements before swapping:", num1, num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("Elements after swapping:", num1, num2)

}
