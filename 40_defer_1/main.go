// Defer in go
// A defer statement defers the execution of a
// function until the surrounding function returns.
package main

import "fmt"

func main() {
  defer fmt.Println("!")
	defer fmt.Println("world") // prints first
	fmt.Println("hello")
}
