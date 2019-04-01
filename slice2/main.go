// Slice in Golang
package main

import (
  "fmt"
)

func main() {
  a := []int{}
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
  a = append(a,1) // slice capacity changes
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
  a = append(a,2,3,4,5) // slice capacity changes
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
  a = append(a,[]int{6,7,8,9}...) // slice capacity changes 
  fmt.Println(a)
  fmt.Printf("length : %v\n ", len(a))
  fmt.Printf("capacity : %v\n ", cap(a))
}
