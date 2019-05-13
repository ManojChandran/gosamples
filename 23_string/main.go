// Immutable strings
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello World")
  s := "hello"
  s[0] = 'H'
  fmt.Println(s) //./main.go:11:8: cannot assign to s[0]
}
