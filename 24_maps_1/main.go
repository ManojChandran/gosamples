// go MAP example
package main

import (
  "fmt"
)

func main() {
  m := make(map[string]int) // construct m := map[key]value{}
  m = map[string]int{ // insert multiple key values
    "A": 1001,
    "B": 2001,
  }
  m["C"] = 3001 // single insert
  delete(m, "C") // delete with key value "B"

  fmt.Println("size ", len(m)) // size of map

  for k,v := range m { // iterate map
    fmt.Println(k,"value = ",v)
  }

  pop, ok := m["B"] // check out the value in map
  if ok {
    fmt.Println("value of B =", pop)
  }
}
