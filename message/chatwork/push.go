package chatwork

import (
	"bytes"
	"errors"
	"strings"
	"text/template"

	"gopkg.in/go-playground/webhooks.v3/github"
)

func PushMsg(pl github.PushPayload) (string, error) {
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

	if len(pl.Commits) == 0 {
		return "", errors.New("Empty Commit detected. Skip.")
	}
	return msg.String(), nil
}
