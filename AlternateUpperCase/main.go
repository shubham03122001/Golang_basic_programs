package main

import "fmt"

func main() {
	a := "shubham"
	answer := ""
	for i := 0; i < len(a); i++ {
		// count:=0
		if i%2 != 0 {
			count := a[i] - 32
			answer += string(count)
		} else {
			answer += string(a[i])
		}
	}
	fmt.Println(answer)

}
