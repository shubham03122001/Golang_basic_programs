package main

import "fmt"

func main() {
	Array := []int{1, 2, 3, 2, 1}
	length := len(Array) / 2
	count := 0
	for i := 0; i <= length; i++ {
		if Array[i] == Array[len(Array)-1-i] {
			count++
		}
	}
	if count == length {
		fmt.Println("PERFECT")
	} else {
		fmt.Println("NOT PERFECT")
	}
}
