package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "l|*e*et|c**o|*de"
	Array := strings.Split(s, "|")
	count := 0
	for i := 0; i < len(Array); i++ {
		if i%2 == 0 {
			for _, value := range Array[i] {
				if strings.Contains(string(value), string("*")) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
