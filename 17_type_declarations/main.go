/* A "type" declaration defines a new "named type"
  that has the same underlying type as an existing
  type
*/
package main

import (
  "fmt"
)

type myfloat64 float64

func main() {
  var v myfloat64 = 16.0
  fmt.Printf("value = %g\n",v)
}
