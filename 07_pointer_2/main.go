// pointers in go
// pointer arithematic is not allowed in go
/* it is perfectly safe for a function to return
  address of a local variable */
package main

import (
  "fmt"
)

func main() {
  var p = pointer()
  fmt.Println(p)
  fmt.Println("pointer() == pointer() ", pointer() == pointer())
}

func pointer() *int  {
  v := 1
  return &v
}
