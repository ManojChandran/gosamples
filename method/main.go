// Method is a function with receiver
package main

import (
  "fmt"
)

func main() {
  g := greeter{
      greeting:"Hello",
      name:"Go",
  }
  g.greet()
  g.greetPointer()
  fmt.Println("new name", g.name)
}

type greeter struct {
  greeting string
  name string
}

func (a greeter) greet()  {
    fmt.Println(a.greeting, a.name)
}
// pointer reciever 
func (a *greeter) greetPointer()  {
    fmt.Println(a.greeting, a.name)
    a.name=" "
}
