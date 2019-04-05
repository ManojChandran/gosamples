// Go routine Channels
// multi directional channel
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
    ch <- 200
    wg.Done()
  }()
  go func ()  {
    fmt.Println("Send to ")
    ch <- 100 // send data to channel
    fmt.Println("Recieved",<-ch)
    wg.Done()
  }()
  wg.Wait()
}
