package main

import "fmt"

func main() {
	S := "Geeks for Geeks"
	count := 0
	for i := len(S) - 1; i >= 0; i-- {
		if string(S[i]) == " " {
			break
		}
		count++

	}
	fmt.Println("Count is:", count)
}
