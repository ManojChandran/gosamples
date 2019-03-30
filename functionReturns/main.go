// go function example
// return value, pointer
package main

import (
  "fmt"
)

func main() {
  s:=sum(1,2,3,4,5)
  fmt.Println("The sum is", *s)

  d,err := divide(5.0,0.0)
  if err !=nil {
    fmt.Println(err)
  } else {
    fmt.Println(d)
  }
}
// go functions can return pointers of local variable
func sum(value ...int) *int {
  result := 0
  for _,v := range value {
    result +=v
  }
  return &result
  // its safe unlike other language
  // go promotes the variable to the heep memmory
}

// return multiple values
func divide(a,b float64) (float64, error){
  if b == 0.0 {
    return a / b, fmt.Errorf("cannot divide by zero")
  }
  return a / b, nil
}
