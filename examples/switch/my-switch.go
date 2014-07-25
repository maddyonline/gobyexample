package main

import "fmt"
import "time"

func main() {
  fmt.Println(time.Now())
  // Testing switch fall-through:
  v := 2
  switch v {
  	case 2 : fmt.Println(v)
  	case 3 : fmt.Println(v)
  }
  switch {
  	case v >= 2: fmt.Println(v)
  	case v <= 2: fmt.Println(v)  // This is like an else clause. Short-Circuiting happens
  }
  // The following two statements are equivalent
  if (v%2 == 0) || (v == 2) {
  	fmt.Println(true)
  }
  switch {
  	case (v%2 == 0), (v == 2): fmt.Println(true)
  }
}
