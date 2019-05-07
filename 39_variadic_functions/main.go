// go functions

package main

import (
  "fmt"
)

func main() {
  fmt.Println("1+2+3+4+5 =",sum(1,2,3,4,5))
}

func sum(vals ...int) int {
  total := 0
  for _, val := range vals {
    total += val
  }
  return total
}
