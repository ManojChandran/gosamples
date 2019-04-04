// panic in go
// A panic typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn’t occur during normal operation, or that we
// aren’t prepared to handle gracefully.
package main

import (
  "fmt"
)

func main() {
  fmt.Println("end")
  // defer will happen prior to panic
  defer fmt.Println("this was defered")
  panic("something bad happened")
  fmt.Println("end")
}
