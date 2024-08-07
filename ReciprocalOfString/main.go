package main

import "fmt"

func main() {
	S := "daDc"
	// Output:
	// zyX
	// Explanation:
	// The reciprocal of the first letter 'a' is 'z'.
	// The reciprocal of the second letter 'b' is 'y'.
	// The reciprocal of the third letter ' ' is ' '.
	// The reciprocal of the last letter 'C' is 'X'.
	for _, ch := range S {
		if ch >= 'a' && ch <= 'z' {
			reciprocalOfCharacter := 'z' - (ch - 'a')
			fmt.Printf("Reciprocal of %c is : %c", ch, reciprocalOfCharacter)
			println()
		} else {
			reciprocalOfCharacter := 'Z' - (ch - 'A')
			fmt.Printf("Reciprocal of %c is : %c", ch, reciprocalOfCharacter)
			println()
		}
	}

}
