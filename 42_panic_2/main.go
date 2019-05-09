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
  fmt.Println("start")
  defer fmt.Println("this was defered")
  panic("something bad happened")
  defer fmt.Println("this was not defered")
  fmt.Println("end")
}
// defer needs to be coded before the panic
// this is to make sure the defer function is
// registered before actual panic happens
