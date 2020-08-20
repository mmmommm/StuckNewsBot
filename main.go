package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/mmmommm/stucknews/repository"

	//slackのままimportするとslack-goのpkgと被ってしまうので名前を他のに変更する
	slackdata "github.com/mmmommm/stucknews/slack"
	// "github.com/slack-go/slack"
	// "github.com/slack-go/slack/slackevents"
)

type Payload struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color     string `json:"color"`
	Title string `json:"title"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postmsg" {
		http.NotFound(w, r)
		return
	}
	repository.Copy()
	urlData := slackdata.Createdata()

	//webhookのurlを渡してある
	webhookurl := os.Getenv("WEBHOOK")

	attachment := Attachment{
		"#FFC0CB",
		urlData,
	}
	payload := Payload{
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	log.Println(string(body))
}

func main() {
	http.HandleFunc("/postmsg", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
					log.Fatal(err)
	}
}
