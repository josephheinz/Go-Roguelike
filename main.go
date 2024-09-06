package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "go-roguelike/utils"
  "go-roguelike/gameMap"
)


func main() {
  colorCodes := map[string]string{"red":"\u001b[31m","green":"\u001b[32m","yellow":"\u001b[33m","blue":"\u001b[34m","purple":"\u001b[35m","teal":"\u001b[36m","white":"\u001b[37m","reset":"\u001b[0m"}
  player := utils.Entity{Name:"player",Stats:utils.Stats{Health:10,MaxHealth:10,Strength:1},Icon:'@',Color:colorCodes["red"],Pos:utils.Position{X:5,Y:5}}
  
  fmt.Println(player)
  /*195884
  for key, value := range colorCodes {
    fmt.Println(value, key)
  }
  */

  testmap := gameMap.Map{Height:10,Width:10}
  // gameMap.DisplayMap( gameMap.GenerateMapAscii( testmap ), player.Pos )
  
  reader := bufio.NewReader(os.Stdin)

  for {

    gameMap.DisplayMap( gameMap.GenerateMapAscii( testmap ), player.Pos )

    fmt.Println("Move (WASD)")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    fmt.Println("You typed "+input)

    gameMap.MovePlayer(input, player.Pos, gameMap.GenerateMapAscii( testmap ))

    if input == "q" {
      fmt.Println("Quit game")
      break
    }

  }

}
