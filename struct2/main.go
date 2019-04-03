// go composition using stuct
package main

import (
  "fmt"
)

type Animal struct{
  Name string
  Origin string
}

type Bird struct{
  Animal
  SpeedKPH float32
  CanFLY bool
}
func main() {
  b := Bird{}
  b.Name = "EMU"
  b.Origin = "Australia"
  b.SpeedKPH = 48
  b.CanFLY = false
  fmt.Println(b) // prints {{EMU Australia} 48 false}

}
