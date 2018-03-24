package chatwork

import (
	cw "github.com/griffin-stewie/go-chatwork"
)

type client struct {
	cw *cw.Client
}

type Config struct {
	ApiKey string
}

func New(config *Config) *client {
	return &client{
		cw.NewClient(config.ApiKey),
	}
}
