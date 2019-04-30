// redeclaring variables
package main

/* variable shadowing occurs when a variable declared
within a certain scope (decision block, method, or inner
class) has the same name as a variable declared in an outer
scope. */
import (
	"fmt"
)

var name string = "manoj" // variable declared
func main() {
	var name string = "chandran"
	fmt.Println("name :", name) // variable declared
}
