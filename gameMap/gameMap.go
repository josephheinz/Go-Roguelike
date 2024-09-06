package gameMap
import (
  "fmt"
  "strings"
  "go-roguelike/utils"
)

type Map struct {
  Height    int
  Width     int
}

type MapCell struct {
  Color    string
  Icon     string
  Pos      utils.Position
}

func GenerateMapAscii(mapInput Map) [][]string {
  
  // initialize ascii characters for the map borders
  const horizontalSideAscii string = "-"
  const verticalSideAscii string = "|"
  const emptySpaceAscii string = "."

  // create an empty 2D array to be filled gridMap
  gridMap := make([][]string, mapInput.Height)
  for arr := range gridMap {
    gridMap[arr] = make([]string, mapInput.Width)
  }

  // iterate through the rows and columns and add edges and empty spaces
  for row := 0; row < mapInput.Height; row++ {

    for col := 0; col < mapInput.Width; col++ {
      
      // if the row is the top or bottom, add the horizontalSideAscii to the array row and column
      if row == 0 || row == mapInput.Height - 1 {
        gridMap[row][col] = horizontalSideAscii

      // if the column is the left or right, add the verticalSideAscii to the array row and column
      }else if col == 0 || col == mapInput.Width - 1 {
        gridMap[row][col] = verticalSideAscii
      }else{
        gridMap[row][col] = emptySpaceAscii
      }

    }

  }  

  return gridMap

}

func DisplayMap(mapInput [][]string, playerPos utils.Position) {
  
  fmt.Print("\033[H\033[2J")
  fmt.Print("\u001b[0m")

  mapInput[playerPos.X][playerPos.Y] = "@"

  for _, row := range mapInput {

    for pos := range row {

      fmt.Print(row[pos])

    }
    fmt.Println()

  }

  mapInput[playerPos.X][playerPos.Y] = "."

}

func MovePlayer (input string, playerPos utils.Position, mapInput [][]string) {
  switch strings.ToLower(input){
    case "w":
      if playerPos.Y > 0 {
        playerPos.Y--
    } 
    case "s":
      if playerPos.Y < len(mapInput) - 2 {
        playerPos.Y++
    }
    case "a":
      if playerPos.X > 0 {
      playerPos.X--
    }
    case "d":
      if playerPos.X < len(mapInput[playerPos.X]) - 2 {
      playerPos.X++
    }
    default:
      fmt.Println("test")
  }
}
