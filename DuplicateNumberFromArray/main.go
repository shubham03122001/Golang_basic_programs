package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 2, 4, 5, 1, 6}

	for i := 0; i < len(array); i++ {
		count := 0
		for j := i; j < len(array); j++ {
			if i != j {
				if array[i] == array[j] {
					count++
				}
			}

		}
		if count >= 1 {
			fmt.Println(array[i])
			//break
		}

	}
}
