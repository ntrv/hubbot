package main

import (
	"net/http"
	"os"
	"time"

	chatwork "github.com/griffin-stewie/go-chatwork"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ntrv/hubbot/github"
	"github.com/ntrv/hubbot/handler"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func PostCWActionFunc(msg string, c echo.Context) error {
	cw := chatwork.NewClient(os.Getenv("API_KEY"))
	res, err := cw.PostRoomMessage(os.Getenv("ROOM_ID"), msg)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	return c.JSON(http.StatusOK, res)
}

func PrintActionFunc(msg string, c echo.Context) error {
	return c.String(http.StatusOK, msg)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(
		github.VerifyMiddleware(
			github.VerifyConfig{Secret: os.Getenv("X_HUB_SECRET")},
		),
	)

	hook := github.NewHook()

	hook.RegisterEvents(handler.Push(PrintActionFunc), gh.PushEvent)
	hook.RegisterEvents(handler.PullRequest(PrintActionFunc), gh.PullRequestEvent)

	e.POST("/", hook.ParsePayloadHandler)

	s := &http.Server{
		Addr:         ":1234",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}
