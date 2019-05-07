/* A program may contain multiple declarations
of the same name, as long as each declarations
are in a different lexical block */
package main

import (
  "fmt"
)

func main() {
  x := "hello"
  for _,x := range x {
    x := x + 'A' - 'a'
    fmt.Printf("%c",x) // prints HELLO
  }
  fmt.Println("\noutside loop",x)
}
