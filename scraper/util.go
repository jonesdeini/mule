package scraper

import (
  "fmt"
  "io/ioutil"
  "net/http"
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
