// go routines
package main

import (
  "fmt"
  "time"
)

func main() {
  var msg = "Hello"
  go func()  {
    fmt.Println(msg) // Print Good Bye
  }()
  msg = "Good bye"
  time.Sleep(100 * time.Microsecond)
}
