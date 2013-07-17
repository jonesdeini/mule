package main

import (
  "github.com/jonesdeini/mule/scraper"
)

func main(){
  /* scraper.ScrapeMatches("http://www.sc2ratings.com/season-info.php?season=spl2&section=Round%206") */
  scraper.ScrapePlayers("http://www.sc2ratings.com/players.php?q=")
}
