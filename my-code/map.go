package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["January"] = 1
  m["March"] = 3
  //v = m["January"]  !! Undefined v
  v := m["January"]
  fmt.Println(v)
  fmt.Println(m)
  delete(m, "March")
  fmt.Println(m)

  valueReturned, isPresent := m["March"]
  fmt.Println(valueReturned) // isPresent can be used to distinguish b/w valueReturned == 0 and non-existant key
  switch isPresent {
    case false: fmt.Println("Sorry, not present!")
    case true: fmt.Println("Yups!")
  }

  // 'static' maps
  m2 := map[string]int {"true" : 0, "false": 1}
  fmt.Println(m2)
}