// go composition using stuct
// tag values in struct and retrive it
package main

import (
  "fmt"
  "reflect"
)

type Animal struct{
  Name string `max:"15"`
  Origin string
}

func main() {
  t := reflect.TypeOf(Animal{})
  field,_ := t.FieldByName("Name")
  fmt.Println(field.Tag)
  fmt.Println(field.Tag.Get("max"))
  fmt.Println("Field count :",t.NumField())
}
