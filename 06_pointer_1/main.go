// pointers in go
// pointer arithematic is not allowed in go
// Variables are sometimes described as addressable values
package main

import (
  "fmt"
)

func main() {
  var a int = 100

  fmt.Println("value of a =",a)
  fmt.Println("address of a =",&a)

  var b *int = &a // b, of type *int, points to a
  fmt.Println("value of b =",*b) // * is used to de-reference
  fmt.Println("address of b =", b)
  a = 200 // change value of a
  fmt.Println("change value of a to 200, value of b =",*b)

  // zero value for a pointer of any type is nil
  var point *int
  fmt.Println("point == nil: ",point) // point == nil

}
