package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/webhooks.v3/github"
	gh "github.com/ntrv/hubbot/github"
)

func (cl client) HandlePullRequest(f PostProcessFunc) gh.ProcessPayloadFunc {
	return func(payload interface{}, c echo.Context) error {
		pl, ok := payload.(github.PullRequestPayload)
		if !ok {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"Failed to parse PRPayload",
			)
		}

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
