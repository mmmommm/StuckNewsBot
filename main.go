package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
    "os"
    "io/ioutil"
    "fmt"

	"github.com/mmmommm/stucknews/repository"

	//slackのままimportするとslack-goのpkgと被ってしまうので名前を他のに変更する
	slackdata "github.com/mmmommm/stucknews/slack"
	// "github.com/slack-go/slack"
	// "github.com/slack-go/slack/slackevents"
)
type Payload struct {
    Text string `json:"text"`
    Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
    Color string `json:"color"`
    TitleLink string `json:"title_link"`
}

func main() {
    // api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
    repository.Copy();
    urlData := slackdata.Createdata()

    //webhookのurlを渡してある
    webhookurl := os.Getenv("WEBHOOK")

    attachment := Attachment {
        "#FFC0CB",
        urlData,
    }
    payload := Payload {
        "今日のニュースだよ！！！！！！",
        []Attachment{attachment},
    }
    //payloadをjsonの形に
    params, err := json.Marshal(payload)
    if err != nil {
        log.Println(err)
        return
    }

    res, err := http.PostForm(
        webhookurl,
        url.Values{"payload": {string(params)}},
    )

    if err != nil {
        log.Println(err)
        return
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Println(err)
        return
    }
    defer res.Body.Close()
    log.Println(string(body))
};