// Command line argument in go
package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("Path ", os.Args[0])
  fmt.Println("Hello ", os.Args[1])
}
