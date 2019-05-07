// go MAP example
package main

import (
  "fmt"
)

func main() {
  companyEmplID := make(map[string]int)
  companyEmplID = map[string]int{
    "Manoj"   : 74672387,
    "Archana" : 23232336,
    "Amaya"   : 65735672,
    "Steve"   : 76576576,
  }

  // map is always passed by reference, changes impact
  fmt.Println("companyEmplID ",companyEmplID)
  sp := companyEmplID
  delete(companyEmplID, "Amaya")
  fmt.Println("sp =",sp)

  // array is a valid key type
  // slice is not a valis key type
  m := map[[3]int]string{}
  fmt.Println("map[[3]int]string{} = ",m)
}
