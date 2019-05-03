package main

import (
  "fmt"
  "os"
  "strconv"

  "github.com/ManojChandran/gosample/17_package_creation/tempconv"
)

func main()  {
  for _,arg := range os.Args[1:]{
    t, err := strconv.ParseFloat(arg, 64)
    if err != nil{
      fmt.Fprintf(os.Stderr, "cf : %v\n", err)
      os.Exit(1)
    }
  }
  f := tempconv.Fahrenhiet(t)
  c := tempconv.Celsius(t)

  fmt.Printf("%s = %s, %s = %s\n",
            f, tempconv.FToC(c), c, tempconv.CToF(c))
}
