package scraper

import (
  "fmt"
  "github.com/moovweb/gokogiri"
)

func ScrapePlayers(url string) {
  pageSource := retrievePageSource(url)

  doc, err := gokogiri.ParseHtml(pageSource)
  errorHandler(err)
  defer doc.Free()
  players, err := doc.Search("//*[contains(@class, 'item-container')]")
  fmt.Println(players)
}
