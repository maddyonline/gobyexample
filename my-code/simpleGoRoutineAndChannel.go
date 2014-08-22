package main

import "fmt"

func main() {
  c := make(chan string)
  go func() {
    c <- "ping"
  }()
  r := <- c
  fmt.Println(r)
}