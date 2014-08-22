package main

import "fmt"

func main() {
 var i int  // initalizes to 0 
 for i < 3 {
    fmt.Println("Hello World!")
    i += 1
  }
  fmt.Println(i)
  for ; i < 4; i++ {
  	fmt.Println("Hello again!")
  }
  for {
  	fmt.Println("Hello -- one final time!")
  	break
  }
  fmt.Println(i)
  for i := 0; i < 1; fmt.Println("Something") {  // local i hides global i.
  	i += 1
  }
  fmt.Println(i) // global i will be visible
}
