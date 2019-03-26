// Go Numbers
package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("add 1+2=", 1+2)
  fmt.Println("sub 1-2=", 1-2)
  fmt.Println("mul 1*2=", 1*2)
  fmt.Println("div 1/2=", 1/2)
  fmt.Println("mod 1%2=", 1%2)

  fmt.Println("div 20/3", 20/3)  // types are not converted, o/p "6"
  fmt.Println("div 20.0/3", 20.0/3)

  fmt.Println("Exponents : ", math.Pow(2.0, 3)) // using Math package

}
