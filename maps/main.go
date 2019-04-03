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
  fmt.Println(companyEmplID)
  // get a particular value
  fmt.Println("Archana",companyEmplID["Archana"])
  // put value to Map
  companyEmplID["Manju"] = 87463747
  fmt.Println("Manju :",companyEmplID["Manju"])
  // delete the value from the Map
  fmt.Println("Delete Steve :")
  delete(companyEmplID, "Steve")
  fmt.Println("Steve :",companyEmplID["Steve"])
  // check out the value in map
  pop, ok := companyEmplID["Steve"]
  fmt.Println("Steve =", pop, ok)
  // cannot deteremine the order of the value returned by map
  fmt.Println(companyEmplID)
  // count of elements in map
  fmt.Println(len(companyEmplID))
  // map is always passed by reference, chnages impact
  fmt.Println("companyEmplID ",companyEmplID)
  sp := companyEmplID
  delete(companyEmplID, "Amaya")
  fmt.Println("sp =",sp)

  // array is a valid key type
  // slice is not a valis key type
  m := map[[3]int]string{}
  fmt.Println("map[[3]int]string{} = ",m)
}
