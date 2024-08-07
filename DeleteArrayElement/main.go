package main

import "fmt"

func main() {
	var number int
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Enter the number you want to delete:")
	fmt.Scan(&number)
	index := -1

	for i := 0; i < len(array); i++ {
		if array[i] == number {
			index = i
			break
		}
	}
	//If the element is found,shift element to left
	if index != -1 {
		for i := index; i < len(array)-1; i++ {
			array[i] = array[i+1]
		}
		//Set the last element to zero
		array[len(array)-1] = 0
	} else {
		fmt.Println("Array element not found")
	}
	answer := make([]int, len(array)-1)
	for i := 0; i < len(array)-1; i++ {
		answer[i] = array[i]
		fmt.Print(answer[i], " ")
	}

}
