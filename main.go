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
    copy()
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
                        links := search()
                        for _, url := range links {
                            link := string(url)
                            fmt.Print(link)
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

func copy() {
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

// func scraping() (links []string) {
//     mainurl := "https://kabutan.jp"
//     url := "https://kabutan.jp/info/accessranking/2_1"
// 	doc, err := goquery.NewDocument(url)
// 	if err != nil {
// 		fmt.Print("url scraping failed")
//     }
//     doc.Find("table.s_news_list tbody tr td a").Each(func(i int, s *goquery.Selection) {
//         if i < 11 {
//             lead, _ := s.Attr("href")
//             link := (mainurl+ lead)
//             links = append(links, link)
//         }
//     })

//     return links
// }

func search() (links []string){
    url := "https://kabutan.jp"
    fileInfos, _ := ioutil.ReadFile("./data/index.html")
    stringReader := strings.NewReader(string(fileInfos))
    doc, err := goquery.NewDocumentFromReader(stringReader)
    if err != nil {
        log.Fatal(err)
    }
    doc.Find("table.s_news_list tbody tr td a").Each(func(i int, s *goquery.Selection) {
        if i < 11 {
            lead, _ := s.Attr("href")
            link := (url+ lead)
            links = append(links, link)
        }
    })
    return links
}