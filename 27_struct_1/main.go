// A struct is a collection of fields.
package main

import (
  "fmt"
)

type person struct { // declaring struct type
  name string
  age int
}

func main() {
  var v person // creating a struct
  v = person{"Bob",20} // initializing a struct
  fmt.Println(v)

  fmt.Println("Name",v.name) // accessing stuct using "."

  p := &v // assiging address of struct
  fmt.Println("Age",p.age) // accessing stuct using pointer
}
