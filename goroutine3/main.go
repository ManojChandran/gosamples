// go routines
package main

import (
  "fmt"
  "sync"
)
// sync main go routine with
// one we have declared
var wg = sync.WaitGroup{}
var counter = 0

func main() {

  for i := 0; i < 10; i++{
    wg.Add(2)
    go sayHello()
    go increment()
  }
  wg.Wait()
}

func sayHello()  {
  fmt.Printf("Hello %v\n", counter) // Print Good Bye
  wg.Done() // reports when routine done
}

func increment()  {
  counter++ // Print Good Bye
  wg.Done() // reports when routine done
}
