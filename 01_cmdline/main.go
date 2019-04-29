// Command line argument in go
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  // $ go run main.go manoj chandran
  fmt.Println("os.Args[0] ", os.Args[0])
  fmt.Println("os.Args[1] ", os.Args[1])
  fmt.Println("os.Args[1:] ", strings.Join(os.Args[1:]," "))
}
