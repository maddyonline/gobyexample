package main

import "fmt"

func swap (a *int, b *int) {
  t := *a
  *a = *b
  *b = t
}

func main() {
  var a, b int = 2, 3
  fmt.Println(a, b)
  swap(&a, &b)
  fmt.Println(a, b)
}