package main

import "fmt"

type printable interface {
  to_string() string
}



func myFmtPrintln(args ...printable) {
  for  _, arg := range args {
    fmt.Println(arg.to_string())
  }
}

func main() {
  fmt.Println("Haalo!")
}