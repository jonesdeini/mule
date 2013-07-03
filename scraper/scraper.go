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
  matchWrappers, err := doc.Search(".//*[@class='match-wrapper']")
  errorHandler(err)

  fmt.Println(parseMatchWrappers(matchWrappers))
}

func retrievePageSource(url string) []byte {
  resp, err := http.Get(url)
  errorHandler(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  return body
}

func parseMatchWrappers(matchWrappers []xml.Node) []match {
  out := []match{}
  for i := range matchWrappers {
    matchWrapperInfo, err := matchWrappers[i].Search(".//*[@class='info']")
    errorHandler(err)
    for j := range matchWrapperInfo {
      res, err := matchWrapperInfo[j].Search(".//dd")
      errorHandler(err)
      if len(res) == 5 {
        temp := match{res[4].Content(), res[2].Content()}
        out = append(out, temp)
      }
    }
  }
  return out
}
