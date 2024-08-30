package main
import (
  "fmt"
  "go-roguelike/utils"
)


func main() {
  colorCodes := map[string]string{"red":"\u001b[31m","green":"\u001b[32m","yellow":"\u001b[33m","blue":"\u001b[34m","purple":"\u001b[35m","teal":"\u001b[36m","white":"\u001b[37m","reset":"\u001b[0m"}
  player := utils.Entity{Name:"player",Stats:utils.Stats{Health:10,MaxHealth:10,Strength:1},Icon:"@",Color:colorCodes["red"]}
  
  fmt.Println(player)
  for key, value := range colorCodes {
    fmt.Println(value, key)
  }
}
