package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	servertiming "github.com/mitchellh/go-server-timing"
	gh "github.com/ntrv/hubbot/github"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func PullRequest(f PostProcessFunc) gh.ProcessPayloadFunc {
	return func(payload interface{}, c echo.Context) error {
		timing := servertiming.FromContext(c.Request().Context())
		m := timing.NewMetric("pr").WithDesc("PullRequest").Start()
		defer m.Stop()

		pl, ok := payload.(github.PullRequestPayload)
		if !ok {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"Failed to parse PRPayload",
			)
		}

		ctx := context.Background()
		c.SetRequest(c.Request().WithContext(ctx))
		j, err := json.Marshal(pl)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				"PRPayload is not JSON format",
			)
		}

		return f(string(j), c)
	}
}
