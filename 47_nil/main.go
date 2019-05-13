// "nil" is not nil
package main

import (
  "fmt"
  "os"
)

func foo() error {
  var err *os.PathError = nil
  return err
}

func main() {
  err := foo()
  fmt.Println(err)        // <nil>
  fmt.Println(err == nil) // false
}
