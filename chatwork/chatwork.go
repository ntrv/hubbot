package chatwork

import (
	cw "github.com/griffin-stewie/go-chatwork"
)

type client struct {
	cw *cw.Client
	roomId string
}

type Config struct {
	ApiKey string
	RoomId string
}

func New(config *Config) *client {
	return &client{
		cw.NewClient(config.ApiKey),
		config.RoomId,
	}
}
