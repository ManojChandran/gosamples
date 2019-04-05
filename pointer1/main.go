// pointers in go
// pointer arithematic is not allowed in go
package main

import (
  "fmt"
)

func main() {
  var ms *myStruct
  fmt.Println(ms) // initial value is <nil>
  ms = &myStruct{foo: 42}
  fmt.Println(ms)
  ms1 := new(myStruct)
  //(*ms1).foo = 42 not required compiler understands
  ms1.foo = 42
  fmt.Println(ms1.foo)
}

type myStruct struct {
  foo int
}
