package main

import "fmt"
import "errors"
import "math"

type myError struct {
  val int
  errorCause string
}

func sqrt (x float64) (float64, error) {
  val := 0.0
  err := errors.New("No error")
  if x < 0 {
    err = errors.New("Something went wrong!")
  } else {
    val = math.Sqrt(x)
    err = nil
  }
  return val, err
}

func (errorObj *myError) Error() string {
  return errorObj.errorCause
}

func main() {
  fmt.Println("hey!")
  //x := -1.0
  x := 16.0
  if sqrt_x, err := sqrt(x); err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("Square root of %f is %f\n", x, sqrt_x)
  }

}