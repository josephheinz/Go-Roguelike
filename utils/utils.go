package utils

type Stats struct {
  Health    int
  MaxHealth int
  Strength  int
}

type Position struct {
  X   int
  Y   int
}

type Entity struct {
  Name    string
  Stats   Stats
  Pos     Position
  Icon    string
  Color   string 
}
