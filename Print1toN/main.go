package main

import (
	"fmt"
)

func main() {
	printNumbers(1)
}

func printNumbers(n int) {
	if n > 5 {
		return
	}
	fmt.Println(n)
	printNumbers(n + 1)
}
