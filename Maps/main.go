package main

import "fmt"

func main() {
	FruitShop := make(map[string]int)
	//fmt.Println(makeMap)
	FruitShop["Mango"] = 30
	FruitShop["Apple"] = 60
	FruitShop["Guava"] = 10

	// if value, ok := makeMap["Feb"]; !ok {
	// 	fmt.Println("Value in boolean form", value)
	// 	fmt.Println(ok)
	// }
	fmt.Println(FruitShop)

	for key, value := range FruitShop {
		fmt.Printf("Key is : %s And value is :%d\n", key, value)
	}

	// array := []int{1, 2, 3, 5, 4, 5, 10, 1}

	// arrayMap := make(map[int]int)

	// for _, value := range array {
	// 	arrayMap[value]++
	// }

	//fmt.Println(arrayMap)

	//Count of array without using lem method
	// count := 0
	// for _, value := range arrayMap {
	// 	count += value
	// }
	// fmt.Println(count)

}
