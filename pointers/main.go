// pointers in go
// pointer arithematic is not allowed in go
package main

import (
  "fmt"
)

func main() {
  var a int = 100

  fmt.Println("value of a =",a)
  fmt.Println("address of a =",&a)

  var b *int = &a
  fmt.Println("value of b =",*b) // * is used to de-reference
  fmt.Println("address of b =", b)
  a = 200 // change value of a
  fmt.Println("change value of a to 200, value of b =",*b)



}
