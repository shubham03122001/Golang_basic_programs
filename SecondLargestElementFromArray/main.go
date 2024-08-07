package main

import "fmt"

func main() {
	array := []int{12, 3, 4, 11, 12}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}

	max := array[0]
	element := array[len(array)-1]

	for i := 0; i < len(array); i++ {
		if array[i] > element && array[i] < max {
			element = array[i]
		}
	}
	fmt.Println("Second Max:", element)
}
