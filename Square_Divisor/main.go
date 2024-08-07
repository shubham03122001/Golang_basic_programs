package main

import "fmt"

func main() {
	// 	Input: 36
	// Output: 4
	// Explaination: The factors which are perfect
	// square are 1, 4, 9 and 36.
	Number := 36
	count := 0
	for i := 1; i <= 36; i++ {
		Square := i * i

		if Number%Square == 0 {
			count++
		}
	}
	fmt.Println(count)

}
