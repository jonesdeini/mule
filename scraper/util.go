package scraper

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "regexp"
)

func errorHandler(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func retrievePageSource(url string) []byte {
  resp, err := http.Get(url)
  errorHandler(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  return body
}

func parseRace(html string) string {
  regex, err := regexp.Compile(`Protoss|Random|Terran|Zerg`)
  errorHandler(err)
  return regex.FindString(html)
}
