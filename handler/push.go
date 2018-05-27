package handler

import (
	"net/http"

	"github.com/labstack/echo"
	gh "github.com/ntrv/hubbot/github"
	"gopkg.in/go-playground/webhooks.v3/github"
	"github.com/ntrv/hubbot/message/chatwork"
)

func Push(f PostProcessFunc) gh.ProcessPayloadFunc {
	return func(payload interface{}, c echo.Context) error {
		pl, ok := payload.(github.PushPayload)
		if !ok {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"Failed to parse PushPayload",
			)
		}

		msg, err := chatwork.PushMsg(pl)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return f(msg, c)
	}
}
