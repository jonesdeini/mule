package scraper

import(
  "encoding/json"
)

type player struct {
  Leagues []string
  RealName string
  Races []string
  Tags string
  /* tags []string */
}

func marshalSlice(players []player) []byte {
  out := []byte{}
  for i := range players {
    playerJson, err := json.Marshal(players[i])
    errorHandler(err)
    out = append(out, playerJson...)
  }
  return out
}
