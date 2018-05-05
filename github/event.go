package github

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

type ProcessPayloadFunc func(payload interface{}, c echo.Context) error

type Webhook struct {
	eventName  gh.Event
	eventFuncs map[gh.Event]ProcessPayloadFunc
}

func NewHook() Webhook {
	return Webhook{
		eventName:  gh.Event(""),
		eventFuncs: map[gh.Event]ProcessPayloadFunc{},
	}
}

func (hook Webhook) RegisterEvents(fn ProcessPayloadFunc, events ...gh.Event) {
	for _, event := range events {
		hook.eventFuncs[event] = fn
	}
}

func (hook Webhook) ParsePayloadHandler(c echo.Context) error {
	if c.Request().Method != echo.POST {
		return echo.NewHTTPError(
			http.StatusMethodNotAllowed,
			fmt.Sprintf(
				"Attempt made using following method is not allowed: %s",
				c.Request().Method,
			),
		)
	}

	event := c.Request().Header.Get("X-GitHub-Event")
	if len(event) == 0 {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"Missing X-GitHub-Event Header",
		)
	}

	hook.eventName = gh.Event(event)
	fn, ok := hook.eventFuncs[hook.eventName]
	if !ok {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf(
				"Webhook Event %s not registered, it is recommended to setup only events in github that will be registered in the webhook to avoid unnecessary traffic and reduce potential attack vectors.",
				event,
			),
		)
	}

	payload, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Cannot read payload",
		)
	}
	return hook.runProcess(payload, fn, c)
}
