// redeclaring variables
package main

import (
  "fmt"
)
var name string = "manoj"
func main() {
  fmt.Println("name :",name)
  // name declaration inside the fucntion
  // took precedence
  var name string = "chandran"
  fmt.Println("name :", name)
}
