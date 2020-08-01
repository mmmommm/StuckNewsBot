package slack

import (
	"fmt"
	"strings"

	"github.com/mmmommm/stucknews/repository"
)

type SlackRepository interface {
	Post(path string, msg []*Post) error
}

type slackImpl struct {
}

// func NewSlackRepository() SlackRepository {
// 	return &slackImpl{}
// }

//slackに投稿する中身
type Post struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Link string `json:"link"`
}

type notifyImpl struct {
}

	// post := []*Post{}
	// line := n.createData("今日のカブタンnews", "#4286f4", link)
	// post = append(post, line)

//repositoty/scraping.goからlinksをとってくる
var links = repository.Scraping()
func (n *notifyImpl) CreateData(title string, color string, links []string) *Post {
	text := []string{}
	//linksをforで回してlinkに入れる
	for i, link := range links {
		//10件までしか取れないようにする
		if i >= 10 {
			break
		}
		//[1]　newsのURLのような形にする
		text = append(text, fmt.Sprintf("[%d] <%s>", i+1, link))
	}
	post := &Post {
		Title: title,
		Color: color,
		Link: strings.Join(text, "\n"),
	}
	return post
}
//このpostの中身をslackで送信したい