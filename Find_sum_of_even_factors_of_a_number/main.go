package main

import "fmt"

func main() {
	// 	Input:
	// N = 30
	// Output:
	// 48
	// Explanation:
	// Even dividers sum 2 + 6 + 10 + 30 = 48

	Number := 30
	Sum := 0
	for i := 2; i <= Number; i += 2 {
		if Number%i == 0 {
			Sum += i
		}
	}
	fmt.Println(Sum)

}
