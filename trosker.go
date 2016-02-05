package kanal

import "time"

type Trosker interface {
	// Auth takes pubkey and seckey from kanal.toml and authenicates
	// the third-party api.
	Auth()
	Get(q Query) *Response
	Timeout(t time.Time) bool
}
