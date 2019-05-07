// Defer in go
// A defer statement defers the execution of a
// function until the surrounding function returns.
package main

import "fmt"

func main() {
  // defer is going to take the
  // value at the time of calling
  a := "start"
  defer fmt.Println(a)
  a = "end"
}
