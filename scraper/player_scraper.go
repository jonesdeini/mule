package scraper

import (
  "fmt"
  "github.com/moovweb/gokogiri"
  "github.com/moovweb/gokogiri/xml"
)

func ScrapePlayers(url string) {
  pageSource := retrievePageSource(url)

  doc, err := gokogiri.ParseHtml(pageSource)
  errorHandler(err)
  defer doc.Free()
  players, err := doc.Search("//*[contains(@class, 'item-container')]")
  parsePlayers(players)
  /* results := parsePlayers(players) */
  /* fmt.Println(results[0]) */
}

func parsePlayers(players []xml.Node) []player {
  out := []player{}

  /* fmt.Println(parseLeagues(players[0]) */

  for i := range players {
    temp := player{}
    itemContainer := players[i].FirstChild()
    temp.realName = itemContainer.FirstChild().Content()
    temp.tags = itemContainer.NextSibling().Content()

    // parse races
    races, err := players[i].Search(".//img")
    errorHandler(err)
    for j := range races {
      temp.races = append(temp.races, parseRace(races[j].String()))

    }

    out = append(out, temp)

  }
  return out
}

/*
  Not working because:
  http://www.sc2ratings.com/players.php?realname=Yang,%20Hee-Soo
  is parsed as:
  http://www.sc2ratings.com/players.php?realname=Yang, Hee-Soo
*/
func parseLeagues(player xml.Node) []string {
  out := []string{}
  partialUrl, err := player.Search(".//a/@href")
  errorHandler(err)
  if len(partialUrl) == 1 {
    playerPageUrl := "http://www.sc2ratings.com/" + partialUrl[0].String()
    playerPageSource  := retrievePageSource(playerPageUrl)

    playerPage, err := gokogiri.ParseHtml(playerPageSource)
    errorHandler(err)
    defer playerPage.Free()
    fmt.Println(playerPage)
  }
  return out
}
