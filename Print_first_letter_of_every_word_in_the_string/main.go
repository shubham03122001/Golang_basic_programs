package main

import (
	"fmt"
	"strings"
)

func main() {
	//Print first letter of every word in the string

	// 	S = "geeks for geeks"
	// Output: gfg

	String := "bad is good"
	Answer := ""

	Array := strings.Split(String, " ")

	for _, Word := range Array {
		Answer += string(Word[0])
	}
	fmt.Println(Answer)

}
