package main

import "fmt"

func main() {
	//Input: S = "$Gee*k;s..fo, r'Ge^eks?"
	//Output: GeeksforGeeks
	String := "$Gee*k;s..fo, r'Ge^eks?"
	Answer := ""
	for i := 0; i < len(String); i++ {
		IntegerValue := String[i]

		if IntegerValue >= 65 && IntegerValue <= 90 || IntegerValue >= 97 && IntegerValue <= 122 {
			Answer += string(IntegerValue)

		}
	}
	fmt.Println(Answer)

}
