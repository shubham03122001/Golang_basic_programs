package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's Football" //Input
	//Output: "s'teL llabtoof"
	result := reverseWords(s)
	fmt.Println(result)

}

func reverse(s string) string {
	n := len(s)
	reverse := make([]byte, n)
	for i := 0; i < n; i++ {
		reverse[i] = s[n-1-i]
	}
	return string(reverse)
}

func reverseWords(s string) string {
	//Split the input string by space
	words := strings.Split(s, " ")

	//Reverse each word and store back it to array/slice
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}
	//Join the words with space and return the result
	return strings.Join(words, " ")

}
