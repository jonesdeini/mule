package scraper

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "github.com/moovweb/gokogiri"
  "github.com/moovweb/gokogiri/xml"
)

func ScrapeAllTheThings(url string) {
  pageSource := retrievePageSource("http://www.sc2ratings.com/season-info.php?season=spl2&section=Round%206")

  doc, err := gokogiri.ParseHtml(pageSource)
  errorHandler(err)
  defer doc.Free()
  highest_container, err := doc.Search(".//*[@class='season-round-date-container']")
  errorHandler(err)

  scrape_children_of_highest_container(highest_container)
}

func retrievePageSource(url string) []byte {
  resp, err := http.Get(url)
  errorHandler(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  return body
}

func scrape_children_of_highest_container(highest_container []xml.Node) {
  for i := range highest_container {
    headline, err := highest_container[i].Search(".//*[@class='headline']")
    errorHandler(err)

    subhead, err := highest_container[i].Search(".//*[@class='sub-head']")
    errorHandler(err)

    fmt.Println(headline)
    fmt.Println(subhead)

    matchWrapper, err := highest_container[i].Search(".//*[@class='match-wrapper']")
    errorHandler(err)
    for j := range matchWrapper {
      playerName, err := matchWrapper[j].Search(".//*[@class='player-link']")
      errorHandler(err)
      fmt.Println(playerName)
    }
  }
}
