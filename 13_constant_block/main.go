// Go Constants
package main

import (
  "fmt"
)

const (
  a = 1
  b
  c = 2
  d
)

func main() {
  fmt.Println("value :", a,b,c,d) // prints 1 1 2 2
}
