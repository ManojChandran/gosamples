/* a closure is a technique for implementing
a lexically scoped name binding in a language
with first-class functions - Wikipedia */

package main

import (
  "fmt"
)

func incrementer() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  increment := incrementer()
  fmt.Println(increment())
  fmt.Println(increment())
  fmt.Println(increment())
  //re assign
  increment = incrementer()
  fmt.Println(increment())


}
