package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ntrv/hubbot/chatwork"
	"github.com/ntrv/hubbot/github"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.NewHook()
	cw := chatwork.New(&chatwork.Config{ApiKey: "hogehoge"})

	hook.RegisterEvents(cw.HandlePush, gh.PushEvent)
	hook.RegisterEvents(cw.HandlePullRequest, gh.HandlePullRequest)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(
		github.VerifyMiddleware(
			&github.VerifyConfig{secret: "hogehoge"},
		),
	)

	e.POST("/", hook.ParsePayloadHandler)

	e.Logger.Fatal(e.Start(":1234"))
}
