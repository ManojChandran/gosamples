package main

import (
  "fmt"
)

func main() {
  first, second := "first", "second"
  fmt.Println(first,":", second)
  // tuple assignment
  first, second = second, first // swapping
  fmt.Println(first,":", second)
}
