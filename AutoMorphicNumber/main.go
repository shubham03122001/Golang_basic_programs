package main

import (
	"fmt"
	"strconv"
)

func main() {
	//An automorphic number is an integer whose square ends with the given integer,
	// as (25)2 = 625, and (76)2 = 5776. Strobogrammatic numbers read the same after having been rotated.
	Number := 76
	Count := 0
	// We have to assign the value of Number variable to another variable as value of Number is going to change
	Value := Number
	//NewNumber := 0

	for Number != 0 {
		Number /= 10
		Count++
	}
	fmt.Println("Count of digit is:", Count)

	Square := Value * Value
	fmt.Println("Square of number is", Square)
	String := strconv.Itoa(Square)
	fmt.Println("After converting the int datatype to string:", String)
	Answer := ""
	if Value > 25 {
		for i := len(String) - 1; i >= Count; i-- {
			Answer = string(String[i]) + Answer
		}
	} else {
		if Count > 1 {
			for i := len(String) - 1; i >= Count-1; i-- {
				Answer = string(String[i]) + Answer
			}
		} else {
			for i := len(String) - 1; i >= Count; i-- {
				Answer = string(String[i]) + Answer
			}
		}
	}
	NewNumber, _ := strconv.Atoi(Answer)
	fmt.Println("After truncating the digit of number:", NewNumber)

	if Value == NewNumber {
		fmt.Println("YES it's AutoMorphicNumber")
	} else {
		fmt.Println("NO")
	}

}
