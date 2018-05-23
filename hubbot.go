package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ntrv/hubbot/chatwork"
	"github.com/ntrv/hubbot/github"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.NewHook()
	cw := chatwork.New(
		&chatwork.Config{
			ApiKey: os.Getenv("API_KEY"),
			RoomId: os.Getenv("ROOM_ID"),
		},
	)

	hook.RegisterEvents(cw.HandlePush, gh.PushEvent)
	hook.RegisterEvents(cw.HandlePullRequest, gh.PullRequestEvent)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(
		github.VerifyMiddleware(
			github.VerifyConfig{Secret: "hogehoge"},
		),
	)

	e.POST("/", hook.ParsePayloadHandler)

	s := &http.Server{
		Addr:         ":1234",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}
