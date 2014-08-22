package main

import "fmt"

func intSeq() func () int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func shiftedIntSeq() func(int) int {
  i := 0
  return func(shift int) int {
    i += 1
    return i + shift
  }
}

func main() {
  intseq := intSeq()
  fmt.Println(intseq())
  fmt.Println(intseq())
  fmt.Println(intseq())
  fmt.Println(intSeq()())

  fmt.Println("Crazy")
  
  fmt.Println(shiftedIntSeq()(2))
  fmt.Println(shiftedIntSeq()(2))
  fmt.Println(shiftedIntSeq()(3))
}