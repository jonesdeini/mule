package scraper

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "github.com/moovweb/gokogiri"
  "github.com/moovweb/gokogiri/xml"
)

func ScrapeAllTheThings(url string) {
  pageSource := retrievePageSource(url)

  doc, err := gokogiri.ParseHtml(pageSource)
  errorHandler(err)
  defer doc.Free()
  matches, err := doc.Search(".//*[@class='item-container clearfix match collapsed']")
  errorHandler(err)

  fmt.Println(parseMatches(matches))
}

func retrievePageSource(url string) []byte {
  resp, err := http.Get(url)
  errorHandler(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  return body
}

func parseMatches(matches []xml.Node) []match {
  out := []match{}
  for i := range matches {
    temp := match{}

    res, err := matches[i].Search(".//*[@class='spoilers winner']")
    errorHandler(err)
    if len(res) == 1 {
      winrar, err := res[0].Search(".//*[@class='player-name']")
      errorHandler(err)
      temp.winrar = winrar[0].Content()
    }

    res, err = matches[i].Search(".//dd")
    errorHandler(err)
    if len(res) == 5 {
      temp.matchDate = res[4].Content()
      temp.teams = res[2].Content()
    }
    out = append(out, temp)
  }
  return out
}
