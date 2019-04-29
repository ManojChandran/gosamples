package main

import (
  "fmt"
)

func init()  {
 fmt.Println("init(), function execut before Main !!!!")
}

func main() {
  fmt.Println("Main function execute after init() !!!")
}
