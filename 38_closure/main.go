/* a closure is a technique for implementing
a lexically scoped name binding in a language
with first-class functions - Wikipedia */

package main

import (
  "fmt"
)

func main() {
  f := squares()
  fmt.Println(f())
  fmt.Println(f())
  fmt.Println(f())
  //re assign
  f = squares()
  fmt.Println(f())
}

func squares() func() int {
  var i int
  return func() int {
    i++
    return i*i
  }
}
