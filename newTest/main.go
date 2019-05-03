/*
  *Read input from STDIN. Print your output to STDOUT
  *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
*/

package main
import (
    "fmt"
    "sort"
    "os"
    "bufio"
    "strings"
  )

func main() {
  gameChance := "WIN"
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter testNumber: ")
  testNumber, _ := reader.ReadString('\n')

  fmt.Print("Enter testMemberCount: ")
  testMemberCount, _ := reader.ReadString('\n')

  var testVillianStrength []int
  fmt.Print("Enter testVillianStrength: ")
  villanStrength, _ := reader.ReadString('\n')
  
  var testPlayerStrength []int
  fmt.Print("Enter testPlayerStrength: ")
  playerStrength, _ := reader.ReadString('\n')

  sort.Ints(testVillianStrength)
  sort.Ints(testPlayerStrength)
  fmt.Println(testNumber)
  fmt.Println(testMemberCount)
  fmt.Println(testVillianStrength)
  fmt.Println(testPlayerStrength)
  for i:=0; i < testMemberCount; i++  {
    if testVillianStrength[i] > testPlayerStrength[i]{
      gameChance = "LOSE"
      break;
    }
  }
  fmt.Println(gameChance)
}
