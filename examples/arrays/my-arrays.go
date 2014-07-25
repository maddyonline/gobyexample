package main

import "fmt"

func main(){
	var chars [26]string

	for i := 0; i < 26; i++ {
		chars[i] = string('a' + i)		// assignment to an array
	}
	fmt.Println(chars)

	var char_counts [26]int				// automatically initialized to 0
	
	const test_string = "teststring"
	for i := 0; i < len(test_string); i++ {
		char_counts[test_string[i] - 'a'] += 1
	}
	fmt.Println(char_counts)

	// 2D arrays
	var matrix [3][2]int // I'm thinking 3 rows and 2 cols
	fmt.Println(matrix)
}

