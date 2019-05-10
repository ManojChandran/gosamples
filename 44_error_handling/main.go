// go error handling
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("sample error message")
	if err != nil {
		fmt.Println(err)
	}
}
