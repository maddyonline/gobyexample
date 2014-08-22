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
}
