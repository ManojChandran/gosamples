// redeclaring variables
package main
/* variable shadowing occurs when a variable declared 
within a certain scope (decision block, method, or inner 
class) has the same name as a variable declared in an outer 
scope. */
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
  // Convert variable type
  var i int = 42
  fmt.Printf("%v, %T\n",i, i)
  var j string
  j = string(i)
  fmt.Printf("%v, %T\n",j, j)
}
