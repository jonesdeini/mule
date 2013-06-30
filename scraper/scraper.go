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
  highest_container, err := doc.Search(".//*[@class='season-round-date-container']")
  errorHandler(err)

  fmt.Println(scrape_children_of_highest_container(highest_container))
}

func retrievePageSource(url string) []byte {
  resp, err := http.Get(url)
  errorHandler(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  return body
}

func scrape_children_of_highest_container(highest_container []xml.Node) []match {
  out := []match{}
  for i := range highest_container {
    headline, err := highest_container[i].Search(".//*[@class='headline']")
    errorHandler(err)

    subhead, err := highest_container[i].Search(".//*[@class='sub-head']")
    errorHandler(err)

    for j := range subhead {
      temp := match{matchDate:headline[0].String(), teams:subhead[j].String()}
      out = append(out,temp)
    }

  }
  return out
}
