package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter number:")
	fmt.Scanln(&number)
	var sum int = 1

	for i := number; i > 0; i-- {
		sum = sum * i
	}
	fmt.Println(sum)

}
