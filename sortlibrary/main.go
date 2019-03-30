// sort function in go
package main

import (
  "fmt"
  "sort"
)

func main() {
  people := []struct{
    Name string
    Age int
  }{
    {"Manoj", 35},
    {"Archana", 30},
    {"Amaya", 24},
  }
  sort.Slice(people, func(i int, j int) bool {
    return people[i].Name < people[j].Name
  })
  fmt.Println("By Name : ", people)
  sort.Slice(people, func(i int, j int) bool {
    return people[i].Age < people[j].Age
  })
  fmt.Println("By age : ", people)
}
