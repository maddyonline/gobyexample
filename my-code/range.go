package main

import "fmt"

func main() {
  var arr = []int {1, 2, 3}
  arr2 := []int {4, 5, 6}
  fmt.Println(arr)
  fmt.Println(arr2)

  m := make(map[int]int)
  for i, val := range arr {
    fmt.Printf("arr[%d] = %d\n", i, val)
    m[i] = val
  }
  fmt.Println(m)

  for k, v := range m {
    fmt.Printf("m[%d] = %d\n", k, v)
  }

}