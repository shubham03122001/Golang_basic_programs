package main

import "fmt"

func main() {
	number := 123
	multiplication := 1
	for {
		remainder := number % 10
		multiplication *= remainder
		number /= 10
		if number == 0 {
			break
		}
	}
	fmt.Println("Multiplication of number:", multiplication)

	var a byte = 3
	b := float64(a)
	fmt.Println(b)
}
