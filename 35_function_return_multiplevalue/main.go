// go function example
// return value, pointer
package main

import (
  "fmt"
)

func main() {
  d,err := divide(5.0,0.0)
  if err !=nil {
    fmt.Println(err)
  } else {
    fmt.Println(d)
  }
}

// return multiple values
func divide(a,b float64) (float64, error){
  if b == 0.0 {
    return a / b, fmt.Errorf("cannot divide by zero")
  }
  return a / b, nil
}
