package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// IncomingURL - Get it from here https://slack.com/services/new/incoming-webhook

// Slack struct - payload parameter of json to post.
type Slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func (s *Slack) Send(incomingURL string) string {
	jsonparams, err := json.Marshal(&s)
	resp, err := http.PostForm(
		incomingURL,
		url.Values{"payload": {string(jsonparams)}},
	)
	if err != nil {
		fmt.Println("post Err")
		fmt.Println(incomingURL)
		fmt.Println(err)
	}

	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println(body)
		defer resp.Body.Close()
		return string(body)

	} else {
		fmt.Println(err)
		return "err occer"
	}
}
