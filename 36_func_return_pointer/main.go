// go function example
// return value, pointer
package main

import (
  "fmt"
)

func main() {
  s:=sum(1,2,3,4,5)
  fmt.Println("The sum is", *s)
}

// go functions can return pointers of local variable
func sum(value ...int) *int {
  result := 0
  for _,v := range value {
    result +=v
  }
  return &result
  // its safe unlike other language, pointer will not be lost
  // go promotes the variable to the heep memmory
}
