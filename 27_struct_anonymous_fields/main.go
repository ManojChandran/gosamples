// anonymous stuct fields
package main

import (
  "fmt"
)

type data struct {
  string
  int
  bool
}

func main() {
  sample1 := data {"Manoj", 35, false}
  sample1.bool = true
  fmt.Println(sample1.string,sample1.int,sample1.bool)
}
