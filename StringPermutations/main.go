package main

import "fmt"

func main() {
	a := "abc"

	for i := 0; i < len(a); i++ {
		answer := ""
		for j := i; j < len(a); j++ {
			answer += string(a[j])
			//fmt.Println(answer)
			fmt.Println(answer)
		}
		//fmt.Println(answer)

	}

}
