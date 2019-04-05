// Go routine Channels
// multi directional channel
// Creating buffer in Channel
// Close the channel
package main

import (
  "fmt"
  "sync"
)

var wg = sync.WaitGroup{}

func main() {
  // Channel buffer "50"
  ch := make(chan int, 50)
  wg.Add(2)
  go func (ch <- chan int)  {
     for i := range ch { // recieve channel
         fmt.Println("Recieved :",i)
     }
     wg.Done()
  }(ch)
  go func (ch chan<- int)  {
    fmt.Println("Send to ")
    ch <- 100 // send data to channel
    ch <- 200 // send data to channel
    close(ch) // close the channel
              // ones closed, no re open available
    wg.Done()
  }(ch)
  wg.Wait()
}
