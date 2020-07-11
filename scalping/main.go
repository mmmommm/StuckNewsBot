package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	doc, err := goquery.NewDocument("https://kabutan.jp/")
	if err != nil {
		fmt.Print("url scalping failed")
	}
	res, err := doc.Find(".acrank_top_news1").Html()
	if err != nil {
		fmt.Print("dom get failed")
	}
	ioutil.WriteFile("path/url.html", []byte(res), os.ModePerm)
}