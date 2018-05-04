package chatwork

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func (cw client) HandlePush(payload interface{}, c echo.Context) error {
	pl, ok := payload.(github.PushPayload)
	if !ok {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Failed to parse PushPayload",
		)
	}
	return c.String(http.StatusOK, pl.HeadCommit.TreeID)
}
