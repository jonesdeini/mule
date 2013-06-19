package main

import (
  "github.com/jonesdeini/mule/scraper"
)

func main(){
  scraper.ScrapeAllTheThings("http://www.sc2ratings.com/season-info.php?season=spl2&section=Round%206")
}
