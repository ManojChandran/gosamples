// Go Constants
package main

import (
  "fmt"
)

const (
  error = iota
  cat
  dog
  )

const (
  crow = iota
)

func main() {
  var animal int = cat
  fmt.Printf("animal :%v\n", animal == cat)

  var bird int
  fmt.Printf("bird :%v\n", bird == crow)
}
