// go composition using stuct
// tag values in struct and retrive it
package main

import (
  "fmt"
)

type Movie struct {
  Title string
  Year int `json:"released"`
  Color bool `json:"color, omitempty"`
  Actors []string
}

func main() {
  fmt.Println("hello")
}
