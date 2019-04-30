// Boolean in go
package main

import (
  "fmt"
)

func main() {
  var istrueornot bool
  istrueornot = true
  fmt.Println(istrueornot)
  someBool := 5 < 3
  fmt.Println(someBool)
  inttobool := inttobool(5)
  fmt.Println(inttobool)
}

func inttobool(i int ) bool {
  return i !=0
}
