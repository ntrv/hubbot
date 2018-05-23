package handler

import (
	"github.com/labstack/echo"
)

//go:generate go-assets-builder -p handler -s="/template" -o assets.go template

type PostProcessFunc func(msg string, c echo.Context) error
