package handler

import (
	"net/http"
	"errors"

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

		if pl.Deleted {
			return echo.NewHTTPError(
				http.StatusNotAcceptable,
				errors.New("Detect delete branch action. Skip."),
			)
		}

		if pl.Created && len(pl.Commits) == 0 {
			return echo.NewHTTPError(
				http.StatusNotAcceptable,
				errors.New("Creating empty branch. Skip."),
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
