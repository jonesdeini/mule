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
  highest_container, err := scrape_highest_container(pageSource)
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

func scrape_highest_container(pageSource []byte) ([]xml.Node, error) {
  doc, err := gokogiri.ParseHtml(pageSource)
  errorHandler(err)
  defer doc.Free()
  return doc.Search(".//*[@class='season-round-date-container']")
}

func scrape_children_of_highest_container(highest_container []xml.Node) {
  for i := range highest_container {
    headline, err := highest_container[i].Search(".//*[@class='headline']")
    errorHandler(err)

    subhead, err := highest_container[i].Search(".//*[@class='sub-head']")
    errorHandler(err)

    fmt.Println(headline)
    fmt.Println(subhead)
  }
}
