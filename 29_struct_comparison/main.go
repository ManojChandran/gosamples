// comparing two struct
/* if all fields of a struct are comparable,
   the struct itself is comparable
*/
package main

import (
  "fmt"
)

type Point struct {
  x, y int
}

func main() {
  p := Point{1,2}
  q := Point{2,3}
  r := Point{1,2}
  fmt.Println(p, q, r)
  fmt.Println("p.x == q.x", p.x == q.x)
  fmt.Println("p == r", p == r)
}
