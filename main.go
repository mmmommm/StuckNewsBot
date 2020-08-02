package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

    "github.com/mmmommm/stucknews/repository"
    //slackのままimportするとslack-goのpkgと被ってしまうので名前を他のに変更する
	slackdata "github.com/mmmommm/stucknews/slack"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
    }
}

func main() {
    repository.Copy()
    http.HandleFunc("/slack/events", Handler)
    
    postMessage := slackdata.Createdata()
    channel := "bot開発"
    jsonStr := `{"channel":"` + channel + `","text":"` + postMessage + `"}`
    //http://crossbridge-lab.hatenablog.com/entry/2017/04/26/000310を参考に実装
    req, err := http.NewRequest(
        "POST",
        "https://hooks.slack.com/services/",
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
            // case slackevents.CallbackEvent:
            // innerEvent := eventsAPIEvent.InnerEvent
            // switch event := innerEvent.Data.(type) {
            //     case *slackevents.AppMentionEvent:
            //     message := strings.Split(event.Text, " ")
            //     if len(message) < 2 {
            //         w.WriteHeader(http.StatusBadRequest)
            //         return
            //     }

            //     command := message[1]
            //     switch command {
            //         //テスト
            //         case "ping":
            //             if _, _, err := api.PostMessage(event.Channel, slack.MsgOptionText("pong", false)); err != nil {
            //                 log.Println(err)
            //                 w.WriteHeader(http.StatusInternalServerError)
            //                 return
            //             }
            //         //newsって打ったら返信でURLが送られてくる
            //         case "news":
            //             //linkにurlをいれる
            //             links := repository.Scraping()
            //             defer os.Remove("./data/index.html")
            //             for _, url := range links {
            //                 link := string(url)
            //                 fmt.Print(link)
            //                 if _, _, err := api.PostMessage(event.Channel, slack.MsgOptionText(link, false)); err != nil {
            //                 log.Println(err)
            //                 w.WriteHeader(http.StatusInternalServerError)
            //                 return
            //                 }
            //             }
            //     }
            // }
    log.Println("[INFO] Server listening")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}