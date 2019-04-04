// recover log function in go
package main

import (
  "fmt"
  "log"
)

func main() {
  fmt.Println("START")
  panicker()
  fmt.Println("END")
}
func panicker()  {
  fmt.Println("About to panic")
  defer func(){
    if err := recover(); err != nil{
      log.Println("Error :", err)
    }
  }()
  panic("done panicking")
  fmt.Println("Done panicking")

}
