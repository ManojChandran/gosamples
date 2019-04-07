// Go channel close
package main

import (
  "fmt"
)

func main() {
  // Channel buffer "50"
  ch := make(chan int, 50)
  // channel can be only closed from
  // sending side
  close(ch)
  fmt.Println(<-ch)
}
