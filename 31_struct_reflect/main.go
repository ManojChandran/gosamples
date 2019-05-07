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
  fmt.Println("Struct name :", t.Name())
  field,_ := t.FieldByName("Name")
  fmt.Println("Field tag   :",field.Tag)
  fmt.Println("Tag value   :",field.Tag.Get("max"))
  fmt.Println("Field count :",t.NumField())
}
