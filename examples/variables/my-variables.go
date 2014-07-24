package main

import "fmt"

func main() {
  var a int = 2
  p, q := 2, 3
  r, s := 2, "astring"
  var u, v = 2, 3
  
  var uninitialized int
 
  fmt.Println(uninitialized) 
  fmt.Println(u, v) // You can't comment this. Because if a variable is declared, it must be used!
  fmt.Println(p, q)
  fmt.Println(r, s)
  fmt.Println(a)
}
