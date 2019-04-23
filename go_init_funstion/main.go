package main

import (
  "fmt"
  _ "os"
)

func init()  {
 fmt.Println("init function execut before Main !!!!")
}

func main() {
  fmt.Println("Main function execute !!!")
}
