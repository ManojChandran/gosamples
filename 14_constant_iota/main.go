// iota keyword represents successive integer constants
// 1,2,3...
package main

import (
  "fmt"
)

func main() {
  const (
    a0 = iota
    a1 = iota
    a2 = iota
  )
  fmt.Println("const value ", a0,a1,a2) // 0,1,2

  const (
    b0 = iota
    b1
    b2
  )
  fmt.Println("const value ", b0,b1,b2) // 0,1,2

  const (
    c0 = iota + 1
    _
    c1
    c2
  )
  fmt.Println("const value ", c0,c1,c2) // 1,3,5
}
