package slack

import (
	"encoding/json"
	"fmt"

	"github.com/mmmommm/stucknews/repository"
)

// type SlackRepository interface {
// 	Post(path string, msg []*Post) error
// }

// type slackImpl struct {
// }

// func NewSlackRepository() SlackRepository {
// 	return &slackImpl{}
// }

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
		text = append(text, fmt.Sprintf("<%s>", link))
	}
	msg, _ := json.Marshal(text)
	strmsg := string(msg)
	//main.goにstring型で渡さないといけないので
	//*Post型を[]byte型に
	// msg, _ := json.Marshal(post)
	//[]byte型をstring型に
	// strmsg := string(msg)
	return strmsg
}

//このpostの中身をslackで送信したい
