package main

import "fmt"

func main() {
	v := [5]int{1, 2, 3, 4, 5}
	fmt.Println(v)

	// A slice is a reference to the subarray
	slice1 := v[1:3]
	fmt.Println(slice1)
	slice1[0] = 9
	fmt.Println(slice1)
	fmt.Println(v)

	// If you want a deep copy, try something like this
	slice2 := make([]int, len(v[1:3]))
	fmt.Println(len(slice2))
	copy(slice2, v[1:3])
	fmt.Println(v)
	fmt.Println(slice2)
	slice2[0] = 7
	fmt.Println(slice2)
	fmt.Println(v)

	// Slices can expand
	slice3 := append(slice2, 33)	// Always returns a deeply-copied item
	fmt.Println(slice3)
	slice3[0] = 0
	fmt.Println(slice3)
	fmt.Println(slice2)

	//Consider vector of adj. lists

	var adj [][]int
	fmt.Println(adj)
	fmt.Println(len(adj))
	n := 5 // Number of vertices
	//adj = make([]int, n)			//  Not allowed
	
	adj2 := make([][]int, n)
	fmt.Println(adj2)
	nbr_lengths := [5]int {1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		adj2[i] = make([]int, nbr_lengths[i])
	}
	fmt.Println("adj2", adj2)

	// Here's one way of making adj similar to adj2
	for i:= 0; i < 5; i++ {
		adj = append(adj, make([]int, nbr_lengths[i]))
	}
	fmt.Println("adj", adj)

}