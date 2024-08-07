package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abcdaab"
	answer := ""

	for i := 0; i < len(a); i++ {
		if !strings.Contains(answer, string(a[i])) {
			answer += string(a[i])
		} else {
			break
		}
	}
	fmt.Println(answer)
}
