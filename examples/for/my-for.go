package main

import "fmt"

func main() {
	var i int
	for {
		i += 1
		if i > 3 {
			break
		}
	}
	fmt.Println("Done after iterations:", i-1)

	fmt.Println("new")
	// Testing scope of j
	for j := 2; j <= 5; j++ {
		fmt.Println(j)
	}
	//fmt.Println("after: ", j) // Not allowed

	// for loop as while loop
	fmt.Println("new")
	for i > 0 {
		fmt.Println(i)
		//--i // not allowed
		i -= 1
	}
}