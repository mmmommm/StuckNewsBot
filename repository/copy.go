package repository

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Copy() {
	url := "https://kabutan.jp/info/accessranking/2_1"
	//urlのbodyを取得
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	//urlのbodyをbodyに代入
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//bodyの内容を./data/index.htmlに出力
	ioutil.WriteFile("./data/index.html", body, 0666)
}