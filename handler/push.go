package handler

import (
	"bytes"
	"net/http"
	"strings"
	"text/template"

	"github.com/labstack/echo"
	gh "github.com/ntrv/hubbot/github"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func genPushMsg(pl github.PushPayload) (string, error) {

	f, err := Assets.Open("/push.tmpl")
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(f); err != nil {
		return "", err
	}

	tpl, err := template.New("Push").Parse(strings.Replace(buf.String(), "\t", "", -1))
	if err != nil {
		return "", err
	}

	msg := &bytes.Buffer{}

	if err := tpl.Execute(msg, pl); err != nil {
		return "", err
	}

	return msg.String(), nil
}

func Push(f PostProcessFunc) gh.ProcessPayloadFunc {
	return func(payload interface{}, c echo.Context) error {

		pl, ok := payload.(github.PushPayload)
		if !ok {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"Failed to parse PushPayload",
			)
		}

		msg, err := genPushMsg(pl)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return f(msg, c)
	}
}
