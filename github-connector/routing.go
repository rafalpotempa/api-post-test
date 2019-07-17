package main

import (
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	// you need to generate personal access token at
	// https://github.com/settings/applications#personal-access-tokens
	personalAccessToken = "4ba7dd9d8fe34a90a292e3ee07021f122cc3ebae"
)

type TokenSource struct {
	AccessToken string
}

type config struct {
	URL     string `json:"url`
	Content string `json:"content_type"`
	Secret  string `json:"secret"`
	Ssl     string `json:"insecure_ssl"`
}

type pld struct {
	Name   string   `json:"name"`
	Config config   `json:"config"`
	Events []string `json:"events"`
	Active bool     `json:"active"`
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte("my-secret-key"))
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	log.Println(event)

	switch e := event.(type) {
	case *github.WatchEvent:
		log.Printf(e.GetSender().GetLogin() + " has started watching your repo")
	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient)
	var cfg config = config{
		URL:     "https://onet.pl",
		Content: "json",
		Secret:  "my-secret-key",
		Ssl:     "0"}
	client.NewUploadRequest
}
