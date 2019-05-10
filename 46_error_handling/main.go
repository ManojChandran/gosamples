// Sentinel error
package main

import (
  "fmt"
  "github.com/pkg/errors"
)

func main() {
  err := errors.New("error occured")
  fmt.Printf("%+v\n",err) // "+" prints full stack trace
}

/*
Annotating errors with errors.Wrap
func Write(w io.Write, buf []byte){
  _,err := w.Write(buf)
  return errors.Wrap(err, "write failed")
}
https://www.youtube.com/watch?v=lsBF58Q-DnY
*/
