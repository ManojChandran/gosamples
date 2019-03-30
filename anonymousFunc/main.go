// anonymous Function in go
package main

import (
  "fmt"
)

func main() {
  func ()  {
      fmt.Println("Anonymous Function")
  }()
  // variable of type function
  var fun func() = func ()  {
      fmt.Println("Function as a type")
  }
  fun()

  f := func ()  {
      fmt.Println("Function as a vriable")
  }
  f()
}
