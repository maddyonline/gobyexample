package main

import "fmt"
import "math"

type complex struct {
  re int
  im int
}

func (c *complex) abs() float64 {
  s := c.re * c.re
  s += c.im * c.im
  return math.Sqrt(float64(s))
}

func printSquares() {
  for i := 1; i < 15; i++ {
    fmt.Println(i*i)
  }
}

func main() {

  c := complex{3, 4}
  d := complex{re: 1, im: 4}
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(c.abs())
  fmt.Println(d.abs())
}