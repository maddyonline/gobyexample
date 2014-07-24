package main

import "fmt"

// Looking at various basic types in Go
func main() {
  fmt.Println("Tom" + "&" + "Jerry")
  fmt.Println("23+23  = ", 23.0 + 23)
  fmt.Println("4.0/2.0 = ", 4.0/2.0)
  fmt.Println("4/2 = " , 4/2)
  fmt.Println("4/5 = " , 4/5)
  fmt.Println("4/5.0 = " , 4/5.0)
  fmt.Println(true, false, true && false, true || false, !true)
}
