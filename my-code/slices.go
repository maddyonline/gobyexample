package main

import "fmt"

func main() {
  std_arr := [4]int {1, 2, 3, 4}
  fmt.Println(std_arr)
  a := make([]int, len(std_arr))
  fmt.Println(a) // [0 0 0 0]
  //copy(a, std_arr) !! Error! 
  copy(a, std_arr[0:]) // This works!
  fmt.Println(a)
  tmp := a[1:3]
  fmt.Println(tmp)
  tmp[0] = 0
  fmt.Println(tmp)
  fmt.Println(a) // tmp is a reference to slice of a
}