// go functions

package main

import (
  "fmt"
)

func main() {
  greetings := "Hello"
  name := "Manoj"
  sayMessage(greetings, name)
  changeMessage(&greetings, &name)
  fmt.Println(greetings, name)
}
// pass by value
func sayMessage(greetings, name string){
  fmt.Println(greetings, name)
}
// pass by reference
func changeMessage(greetings, name *string ){
   *greetings = "bye"
}
