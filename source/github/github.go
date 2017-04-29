package github

import (
	"net/http"
	"time"

	"github.com/lanzafame/kanal"
)

type Github struct {
	TimeoutInterval time.Duration
}

func (gh *Github) Auth() *http.Client {
	//	conf := &oauth2.Config{
	//		ClientID:     "Blah", // Need to input values from kanal.toml here
	//		ClientSecret: "SecretBlah",
	//		Endpoint:     github.Endpoint,
	//	}

	return nil
}

func (gh *Github) Get(q kanal.Query) *kanal.Response {
	return nil
}

func (gh *Github) Timeout() bool {
	return false
}
