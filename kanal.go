package kanal

// Access provides the authentication mechanism for the Troskers and Lukas
type Access interface {
	Auth() (interface{}, error)
	RetrieveCreds(map[string]string) (map[string]string, error)
}

//
type Trosker interface {
	Get(q Query) (*Record, error)
	Access
}

type Luka interface {
	Put(r *Record) error
	Access
}

type Mapping struct {
	Trosker
	Luka
}

type Record interface {
}

type Query interface {
}
