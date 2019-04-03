// Handling files
package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  b,err := ioutil.ReadFile("sample.txt")
  if err != nil {
    fmt.Println(err)
  }

  str:=string(b) // converting byte to string
  fmt.Println(str)
}
