package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	servertiming "github.com/mitchellh/go-server-timing"
	gh "github.com/ntrv/hubbot/github"
	"github.com/ntrv/hubbot/message/chatwork"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func Push(f PostProcessFunc) gh.ProcessPayloadFunc {
	return func(payload interface{}, c echo.Context) error {
		timing := servertiming.FromContext(c.Request().Context())

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

		m := timing.NewMetric("template").WithDesc("Generate Message").Start()
		msg, err := chatwork.PushMsg(pl)
		m.Stop()
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}
		return f(msg, c)
	}
}
