// Slice in Go
package main

import (
  "fmt"
)

func main() {
  // declaring Slice
  a := []int{1,2,3}
  b := a // naturally reference type
  fmt.Println("a :",a)
  b[1]=5
  fmt.Printf("length: %v\n", len(a))
  fmt.Printf("capacity: %v\n", cap(a))
  fmt.Println("b :",b) //prints 1,5,3
  // ":" in Slice
  row0 := []int{1,2,3,4,5,6,7,8,9,0}
  row1 := row0[:] //slice of all element
  row2 := row0[3:] //slice from 4th element to end
  row3 := row0[:6] // slice first 6 element
  row4 := row0[3:6] // slice the 4th,5th,6th elements
  row0[5] = 42 // changes 5th element value in all rows
  fmt.Println("row0 :",row0) //prints 1,5,3
  fmt.Println("row1 :",row1)
  fmt.Println("row2 :",row2)
  fmt.Println("row3 :",row3)
  fmt.Println("row4 :",row4)
}
