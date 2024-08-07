package main

import "fmt"

func main() {
	a := "AaAaAaaA"
	//Output should be 0 as we neglect the shift or caps lock i.e small and capital letter means dont increase the counter
	count := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i]+32 != a[i+1] && a[i]-32 != a[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
