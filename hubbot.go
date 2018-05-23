package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ntrv/hubbot/github"
	"github.com/ntrv/hubbot/handler"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

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

	hook.RegisterEvents(
		handler.Push(handler.PrintPostProcess),
		gh.PushEvent,
	)
	hook.RegisterEvents(
		handler.PullRequest(handler.PrintPostProcess),
		gh.PullRequestEvent,
	)

	e.POST("/", hook.ParsePayloadHandler)

	s := &http.Server{
		Addr:         ":1234",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}
