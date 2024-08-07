package main

import "fmt"

func main() {
	Name := "shubham"
	Answer := ""

	for i := 0; i < len(Name); i++ {
		Answer = string(Name[i]) + Answer
	}
	fmt.Println(Answer)
}
