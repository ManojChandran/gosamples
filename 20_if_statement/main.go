//
package main

import (
  "fmt"
)

func main() {
  number := 0.5
  if number > 1 {
    fmt.Println("number is greater than 1")
  } else if number < 100{
    fmt.Println("number is less than 100")
  } else {
    fmt.Println("number is between 1 and 100")
  }
}
