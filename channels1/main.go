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
  // recieve only channel
  go func (ch <- chan int)  {
    i := <- ch // recieve channel
    fmt.Println("Recieved :",i)
    wg.Done()
  }(ch)
  go func (ch chan<- int)  {
    fmt.Println("Send")
    ch <- 100 // send data to channel
    wg.Done()
  }(ch)
  wg.Wait()
}
