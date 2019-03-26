// GO variables
package main

import (
  "fmt"
)

var i int = 42

// club var together
var (
  name string = "manoj"
  age int = 35
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
  fmt.Printf("%v, %T\n", d, d)
  fmt.Printf("%v, %T\n", e, e)
  fmt.Printf("%v, %T\n", f, f)

  //variable declared for the package
  fmt.Printf("%v, %T\n", i, i)
  fmt.Printf("%v, %T\n", name, name)
  fmt.Printf("%v, %T\n", age, age)
}
