package main

import "fmt"

func simpleAdd(a int, b int) int {
  return a + b
}

func add(ints ...int) int {
  var sum int // auto-initialized to 0
  for _, x := range ints {
    sum += x
  }
  return sum
}

func voidFunc() {
  fmt.Println("void")
}

func weird(strings ...string) {
  for _, x := range strings {
    fmt.Printf("%s ", x)
  }
  fmt.Println()
}

func main() {
  a := []int {1, 2, 3, 4}
  fmt.Println(simpleAdd(2,3))
  fmt.Println(add(1,2,3,4))
  fmt.Println(add(a...))
  fmt.Println(add(a[1:3]...))

  voidFunc()
  
  weird("This", "sure", "rocks!")
}