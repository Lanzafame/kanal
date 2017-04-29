package kanal

import (
	"net/http"
	"time"
)

type Source struct {
	Name string
}

type Store struct {
	Name string
}

type Mapping struct {
	Source
	Store
}

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
