// Stings and chars
// every type in go has a default value\
// single line chars ""
// Multiline ``
// rune ''
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello World")
  // Mutliline string print
  fmt.Println(`This is
    a multi-line. \n
    String
    `)
  fmt.Println("\u2272")
  // singlechars "rune"
  fmt.Println('a')
}
