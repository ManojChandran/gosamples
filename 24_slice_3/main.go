// Slice misc
package main

import (
  "fmt"
)

func main() {
  s :=[]int{1,2,3}
  fmt.Printf("inside main function: \t%p \t%v\n",s,s)  //prints 1,2,3

  doubleValue(s)
  fmt.Printf("inside main function: \t%p \t%v\n",s,s)// prints 2,4,6

  growSlice(s)
  fmt.Printf("inside main function: \t%p \t%v\n",s,s) // prints 2,4,6

}

func doubleValue(s []int){
  for k,v := range s{
    s[k]=v*2
  }
  fmt.Printf("inside doub function: \t%p \t%v\n",s,s) // prints 2,4,6

}

func growSlice(s []int)  {
      s = append(s,4,5,6) // reassign the slice pointer
      fmt.Printf("inside grow function: \t%p \t%v\n",s,s)
      // function checks for the slice size. if the slice size
      // is less than the new appended one, slice copies old values
      // respond back getting old values printed in mainfunction.
}
