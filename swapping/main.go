// Simple Hello World program
// inline comments
/* Multi line
   comments */
package main

import (
  "fmt"
)

func main() {
  first, second := "first", "second"
  fmt.Println(first,":", second)
  // swapping 
  first, second = second, first
  fmt.Println(first,":", second)
}
