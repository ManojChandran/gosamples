// go error handling
package main

import (
	"fmt"
)

func main() {
	const name, id = "MANOJ", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
