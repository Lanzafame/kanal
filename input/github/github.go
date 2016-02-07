package github

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Github struct {
}

func Auth() *http.Client {
	conf := &oauth2.Config{
		ClientID:     "Blah", // Need to input values from kanal.toml here
		ClientSecret: "SecretBlah",
		Endpoint:     github.Endpoint,
	}

	return
}
