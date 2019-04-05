// Go routine Channels
package main

import (
  "fmt"
  "sync"
)

var wg = sync.WaitGroup{}

func main() {
  ch := make(chan int)
  wg.Add(2)
  go func ()  {
    i := <- ch // recieve channel
    fmt.Println("Recieved :",i)
    wg.Done()
  }()
  go func ()  {
    fmt.Println("Send")
    ch <- 100 // send data to channel
    wg.Done()
  }()
  wg.Wait()
}
