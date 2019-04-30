/* The expressioin new(T) creates an unnamed variable
of type T, initializes it to the zero value of T, and
returns a address, which is a value of type *T
*/

package main

import (
  "fmt"
)

func main() {
  p := new(int)
  fmt.Println(p)

  var q = pointer()
  fmt.Println(q)
}

func pointer() *int  {
  return new(int)
}
