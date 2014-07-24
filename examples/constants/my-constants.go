package main

import "fmt"

func main() {
  const tmp = 2    // It's ok to declare constants and not use them, but not quite with vars and packages
  const new_line = "\n"
  const new_line_char = '\n'
  fmt.Println(new_line[0]) // You have to use fmt if you import it
  fmt.Println(new_line_char) // You have to use fmt if you import it
}
