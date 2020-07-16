package copy

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Copy() {
	url := "https://kabutan.jp/info/accessranking/2_1"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./data/index.html", body, 0666)
}