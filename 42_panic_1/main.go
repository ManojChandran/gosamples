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
  defer fmt.Println("run before panic") // function executes before panic
  a, b := 1, 0
  ans := a / b //panic: runtime error: integer divide by zero
  fmt.Println(ans)
  fmt.Println("end")
}
