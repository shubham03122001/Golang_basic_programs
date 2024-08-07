package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 1, 3, 4}
	frequency := make(map[int]int)

	for _, value := range array {
		frequency[value]++
	}
	fmt.Print(frequency)
}
