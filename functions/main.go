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
  sum("The sum is", 1,2,3,4,5,6)
}
// pass by value
func sayMessage(greetings, name string){
  fmt.Println(greetings, name)
}
// pass by reference
func changeMessage(greetings, name *string ){
   *greetings = "bye"
}

func sum(msg string, values ...int)  {
  fmt.Println(values)
  result :=0
  for _,v := range values {
    result +=v
  }
  fmt.Println(msg, result)
}
