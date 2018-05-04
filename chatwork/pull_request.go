package chatwork

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func (cw client) HandlePullRequest(payload interface{}, c echo.Context) error {
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
	return c.JSON(http.StatusOK, j)
}
