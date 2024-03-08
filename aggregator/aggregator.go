package aggregator

import "log"

type Config struct {
	ApiKey string
	Proxy  string
}

type Aggregator interface {
	RegisterCallback()
}

type aggr struct {
	Url    string
	method string

	GetCurrentUrl func() string
}

func New(provider string, config *Config) *aggr {
	return &aggr{}
}

func (m *aggr) RegisterCallback() {
	log.Println("oke")

	m.method = ""
}
