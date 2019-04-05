// go routines RWMutex()
package main

import (
  "fmt"
  "runtime"
  "sync"
)
// sync main go routine with
// one we have declared
var wg = sync.WaitGroup{}
var counter = 0
// Infinte number of readers
// only one writer at a time
var n = sync.RWMutex{}

func main() {
  runtime.GOMAXPROCS(100)
  for i := 0; i < 10; i++{
    wg.Add(2)
    go sayHello()
    go increment()
  }
  wg.Wait()
}

func sayHello()  {
  n.RLock()
  fmt.Printf("Hello %v\n", counter) // Print Good Bye
  n.RUnlock()
  wg.Done() // reports when routine done
}

func increment()  {
  n.Lock()
  counter++ // Print Good Bye
  n.Unlock()
  wg.Done() // reports when routine done
}
