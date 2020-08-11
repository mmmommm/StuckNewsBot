package repository

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraping() (links []string) {
	//urlを踏んだ時のために最初につけるurlを変数として定義
	url := "https://kabutan.jp"
	//./data/index.htmlを読み込む
	fileInfos, _ := ioutil.ReadFile("./data/index.html")
	stringReader := strings.NewReader(string(fileInfos))
	//スクレイピングするファイルを指定
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table.s_news_list tbody tr td a").Each(func(_ int, s *goquery.Selection) {
		//leadに上記のタグのhrefの値を順番に代入
		lead, _ := s.Attr("href")
		//リンクを正常に飛ばすためにhttps://kabutan.jpをくっつける
		link := (url + lead)
		//[]string型であるlinksに前からlinkを挿入する
		links = append(links, link)
	})
	return links
}
