package scraping

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

var href string

func Scraping() (href string){
	url := "https://kabutan.jp/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("url scraping failed")
	}
	doc.Find(".acrank_top_news1 a").Each(func(_ int, s *goquery.Selection) {
		href, _ = s.Attr("href")
	})
	return href
}