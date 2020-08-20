package slack

import (
	// "encoding/json"
	"fmt"
	"strings"

	"github.com/mmmommm/stucknews/repository"
)
//slackに投稿する中身

//repositoty/scraping.goからlinksをとってくる
func Createdata() string {
	text := []string{}
	links := repository.Scraping()
	//linksをforで回してlinkに入れる
	for i, link := range links {
		//10件までしか取れないようにする
		if i >= 10 {
			break
		}
		//"[1]　newsのURL" のような形にする
		text = append(text, fmt.Sprintf("[%d] <%s>", i+1, link))
	}
	slackpost := strings.Join(text, "\n")
	return slackpost
}

//このpostの中身をslackで送信したい
