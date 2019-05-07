// go functions
/*
func name (parameter-list) return-list  {
  body
}
*/

package main

import (
  "fmt"
  "math"
)

func main() {
  h := hypot(2,3)
  fmt.Println(h)
}

func hypot(x, y float64) float64 {
  return math.Sqrt(x*x + y*y)
}
