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

type Int struct { i int}
type String struct {s string}

func (num *Int) to_string() string {
  return "sss"
}
func (str *String) to_string() string {
  return str.s
}

func main() {
  a := Int{3}
  b := String{"Test"}
  fmt.Println(a, b)
  myFmtPrintln(a, b)
  fmt.Println("Haalo!")
}