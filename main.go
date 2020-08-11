package main;

import (
    "bytes"
    "fmt"
    "net/http"
	"github.com/mmmommm/stucknews/repository"
	//slackのままimportするとslack-goのpkgと被ってしまうので名前を他のに変更する
	slackdata "github.com/mmmommm/stucknews/slack"

	// "github.com/slack-go/slack"
	// "github.com/slack-go/slack/slackevents"
)
// type Post struct {
// 	Title string `json:"title"`
// 	Color string `json:"color"`
// 	Link string `json:"link"`
// }

func main() {
    // api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
    repository.Copy();
    urlData := slackdata.Createdata()
    // //http://crossbridge-lab.hatenablog.com/entry/2017/04/26/000310を参考に実装
    //channelIDを取得して変数urlに入れる
    //channel := os.Getenv("SLACK_CHANNEL")
    // post := []*Post{}
    // post = append(post, &Post{
	// 	Title: "今日のカブタンnews",
	// 	Color: "#4286f4",
	// 	Link: strings.Join(urlData, "\n"),
    // })

    // if _, _, err := api.PostMessage(channel, slack.MsgOptionText(post, false)); err != nil {
    //     log.Println(err)
    //     return
    // }

    jsonStr := `{"text":"連絡事項です！！！！",
                    "attachments":[
                        "color":"#4286f4",
                        "pretext":"今日のカブタンニュース",
                        "title_link":"` + urlData + `"
                    ]
                }`
    
    req, err := http.NewRequest(
        "POST",
        "https://hooks.slack.com/services/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",
        bytes.NewBuffer([]byte(jsonStr)),
    )
    if err != nil {
        fmt.Print(err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Print(err)
    }

    fmt.Print(resp)
    defer resp.Body.Close()
}