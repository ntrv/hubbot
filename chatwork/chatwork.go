package chatwork

import (
	"github.com/labstack/echo"
)

//go:generate go-assets-builder -p chatwork -s="/template" -o assets.go template

type client struct{}

type PostProcessFunc func(msg string, c echo.Context) error

func New() *client {
	return &client{}
}
