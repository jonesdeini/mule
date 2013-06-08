package main

import (
  "fmt"
  "github.com/jonesdeini/mule/scraper"
)

func main(){
  pageSource := scraper.RetrievePageSource("http://www.sc2ratings.com/season-info.php?season=spl2&section=Round%206")
  fmt.Println(string(pageSource))
}
