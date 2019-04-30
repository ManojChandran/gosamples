// Go Constants
package main

import (
  "fmt"
)
// Shadowing is allowed for constants
const myConst int16 = 27

func main() {
  const myConst int = 14
  // myConst = 43 not allowed, gives error
  fmt.Printf("%v, %T\n", myConst, myConst)
  const plainConst = 13 // allowed
  fmt.Printf("%v, %T\n", plainConst, plainConst)
}
