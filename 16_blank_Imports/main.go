package main

import (
	myfmt "fmt"
	_ "os" /* The blank identifier allows the
						compiler to accept the import and
						call any init functions that can be
						found in the different code files within
						that package. */
)

func main() {
	myfmt.Println("hello")
}
