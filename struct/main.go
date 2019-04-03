// go anonymous stuct
package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func main() {
  fmt.Println(person{"Bob",20})
  fmt.Println(person{name: "Alice", age: 30})
  s := person{name: "Manoj", age: 35}
  fmt.Println(s)

  sp := &s
  fmt.Println(sp.age)
}
