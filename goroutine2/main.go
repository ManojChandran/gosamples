// go routines
package main

import (
  "fmt"
  "sync"
)
// sync main go routine with
// one we have declared
var wg = sync.WaitGroup{}

func main() {
  var msg = "Hello"
  wg.Add(1) //add the below goo routine
  go func()  {
    fmt.Println(msg) // Print Good Bye
    wg.Done() // reports when routine done
  }()
  msg = "Good bye"
  wg.Wait()
}
