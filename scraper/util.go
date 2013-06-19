package scraper

import (
  "fmt"
)

func errorHandler(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
