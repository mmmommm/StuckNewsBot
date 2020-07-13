package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	// "github.com/mmmommm/stucknews/scraping"

	"github.com/PuerkitoBio/goquery"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func main() {
    api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

    http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {
        verifier, err := slack.NewSecretsVerifier(r.Header, os.Getenv("SLACK_SIGNING_SECRET"))
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        bodyReader := io.TeeReader(r.Body, &verifier)
        body, err := ioutil.ReadAll(bodyReader)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        if err := verifier.Ensure(); err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        switch eventsAPIEvent.Type {
            case slackevents.URLVerification:
                var res *slackevents.ChallengeResponse
                if err := json.Unmarshal(body, &res); err != nil {
                    log.Println(err)
                    w.WriteHeader(http.StatusInternalServerError)
                    return
                }
                w.Header().Set("Content-Type", "text/plain")
                if _, err := w.Write([]byte(res.Challenge)); err != nil {
                    log.Println(err)
                    w.WriteHeader(http.StatusInternalServerError)
                    return
                }
            case slackevents.CallbackEvent:
            innerEvent := eventsAPIEvent.InnerEvent
            switch event := innerEvent.Data.(type) {
            case *slackevents.AppMentionEvent:
                message := strings.Split(event.Text, " ")
                if len(message) < 2 {
                    w.WriteHeader(http.StatusBadRequest)
                    return
                }

                command := message[1]
                switch command {
                    //テスト
                    case "ping":
                        if _, _, err := api.PostMessage(event.Channel, slack.MsgOptionText("pong", false)); err != nil {
                            log.Println(err)
                            w.WriteHeader(http.StatusInternalServerError)
                            return
                        }
                    //newsって打ったら返信でURLが送られてくる
                    case "news":
                        //linkにurlをいれる
                        links := scraping()
                        for _, url := range links {
                            link := string(url)
                            if _, _, err := api.PostMessage(event.Channel, slack.MsgOptionText(link, false)); err != nil {
                            log.Println(err)
                            w.WriteHeader(http.StatusInternalServerError)
                            return
                        }
                        }
                }
            }
        }
    })
    log.Println("[INFO] Server listening")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
func scraping() (links []string) {
	url := "https://kabutan.jp"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("url scraping failed")
    }
    doc.Find("div.acrank_top_news1 > table.visited1 > tbody > tr > td > a").Each(func(_ int, s *goquery.Selection) {
        lead, _ := s.Attr("href")
        link := (url + lead)
        links = append(links, link)
    })
    return links
}