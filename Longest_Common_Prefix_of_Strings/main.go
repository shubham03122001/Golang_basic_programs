package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// 	Input: arr[] = ["geeksforgeeks", "geeks", "geek", "geezer"]
	// Output: gee
	// Explanation: "gee" is the longest common prefix in all the given strings.
	Array := []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	//Array := []string{"aab", "aabc", "aa", "aabc"}
	length := len(Array)

	sort.Slice(Array, func(i, j int) bool {
		return len(Array[i]) < len(Array[j])
	})
	fmt.Println(Array)

	count := 0
	Value := Array[0]
	LengthOfFirstString := len(Value)
	for i := 1; i < length; i++ {
		if strings.HasPrefix(Array[i], Value[0:LengthOfFirstString]) {
			count++
			LengthOfFirstString--
		}
	}
	fmt.Println("Matched length of letter", count)
	fmt.Println(Value[0:count])

}
