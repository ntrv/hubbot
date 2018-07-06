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

func hubbotHandler() echo.HandlerFunc {
	hook := github.NewHook()

	if os.Getenv("ROOM_ID") != "" && os.Getenv("API_KEY") != "" {
		hook.RegisterEvents(
			handler.Push(handler.SendChatworkPostProcess),
			gh.PushEvent,
		)
		hook.RegisterEvents(
			handler.PullRequest(handler.SendChatworkPostProcess),
			gh.PullRequestEvent,
		)
	} else {
		hook.RegisterEvents(
			handler.Push(handler.PrintPostProcess),
			gh.PushEvent,
		)
		hook.RegisterEvents(
			handler.PullRequest(handler.PrintPostProcess),
			gh.PullRequestEvent,
		)
	}
	return hook.ParsePayloadHandler
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if os.Getenv("X_HUB_SECRET") != "" {
		e.Use(
			github.VerifyMiddleware(
				github.VerifyConfig{Secret: os.Getenv("X_HUB_SECRET")},
			),
		)
	}

	e.POST("/", hubbotHandler())

	s := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}
