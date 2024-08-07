package main

import (
	"fmt"
	"strings"
)

func main() {
	//Remove vowels from string
	//welcome to geeksforgeeks
	S := "welcome to geeksforgeeks"
	lowerStr := strings.ToLower(S)
	Answer := ""
	for i := 0; i < len(lowerStr); i++ {
		if lowerStr[i] != 'a' && lowerStr[i] != 'e' && lowerStr[i] != 'i' && lowerStr[i] != 'o' && lowerStr[i] != 'u' {
			Answer += string(lowerStr[i])
		}
	}
	fmt.Println(Answer)
}
