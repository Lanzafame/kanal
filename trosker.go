package kanal

import (
	"net/http"
	"time"
)

type Response interface {
}

type Query interface {
}

type Trosker interface {
	// Auth takes pubkey and seckey from kanal.toml and authenicates
	// the third-party api.
	Auth() *http.Client
	Get(q Query) *Response
	Timeout(t time.Time) bool
}
