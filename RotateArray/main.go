package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	answer := make([]int, len(array))
	// fmt.Println("For how many times you want to shift the array elements")
	// var choice int
	// fmt.Scanln(&choice)
	//count := 0

	for i := 0; i < len(array); i++ {
		if i == 0 {
			answer[len(array)-1] = array[i]
		} else {
			answer[i-1] = array[i]
		}
	}

	for i := 0; i < len(answer); i++ {
		fmt.Print(answer[i], " ")
	}
}
