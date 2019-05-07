package main
import (
    "fmt"
    "os"
    "bufio"
  )

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  inputValue, _ := reader.ReadString('\n')

  fmt.Println(inputValue)
}
