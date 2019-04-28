// Array's in go
package main

import (
  "fmt"
)

func main() {
  // Declare an integer array of five elements.
  var array [5] int
  fmt.Println(array)

  // declaring an array initializing specific elements
  names := [3]string{"Manoj", "Archana", "Amaya"}
  fmt.Println(names)

  // Declaring an array with Go calculating size
  grades := [...]int{97,85,93}
  fmt.Println(grades, "len :", len(grades))

  // array's are considered as values, not like
  // other language when you create an Array
  // its getting pointing to that
  a :=[...]int{1,2,3}
  b := a // value gets actually copied
  c := &a // c is pointing
  b[1] = 5
  c[1] = 0
  fmt.Println(a) // prints 1,2,3
  fmt.Println(b) //prints 1,5,3
  fmt.Println(c) //prints 1,0,3

  var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
  fmt.Println(identityMatrix)
}
