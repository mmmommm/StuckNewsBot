package scraping

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraping() (links []string){
	url := "https://kabutan.jp"
	fmt.Print(url)
    fileInfos, _ := ioutil.ReadFile("./data/index.html")
    stringReader := strings.NewReader(string(fileInfos))
    doc, err := goquery.NewDocumentFromReader(stringReader)
    if err != nil {
        log.Fatal(err)
    }
    doc.Find("table.s_news_list tbody tr td a").Each(func(_ int, s *goquery.Selection) {
        lead, _ := s.Attr("href")
        link := (url+ lead)
        links = append(links, link)
    })
    return links
}