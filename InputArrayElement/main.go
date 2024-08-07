package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of array:")
	fmt.Scanln(&size)
	array := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter array element:")
		fmt.Scanln(&array[i])
	}
	for i := 0; i < size; i++ {
		fmt.Print(array[i], " ")
	}
}
