// go anonymous stuct
package main

import (
  "fmt"
)

func main() {
  aPerson := struct{name string}{name: "Manoj Chandran"}
  anotherPerson := aPerson
  anotherPerson.name = "Archana Premachandran"
  fmt.Println(aPerson)
  fmt.Println(anotherPerson)
  // Using pointer
  anotherPerson1 := &aPerson
  fmt.Println(aPerson)
  anotherPerson1.name = "Archana Premachandran"
  fmt.Println(aPerson)
  fmt.Println(anotherPerson1)

}
