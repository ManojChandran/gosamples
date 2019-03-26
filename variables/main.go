// GO variables
package main

import (
  "fmt"
)

func main() {
  // static declaration
  var x int = 5
  var y int = 6
  var sum1 int = x + y
  fmt.Println(sum1)

  // dynamic declaration
  a := 5
  b := 6
  sum2 := a + b
  fmt.Println(sum2)

  // mixed declaration
  var d, e, f = 3, 4, "foo"
  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f)
}