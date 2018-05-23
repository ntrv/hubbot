package chatwork

import (
	"bytes"
	"net/http"
	"strings"
	"text/template"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func genPushMsg(pl github.PushPayload) (string, error) {

	templateText := `[info][title]Push event at {{.Repository.Name}}[/title]URL: {{.Repository.URL}}
		Ref: {{.Ref}}
		CompareURL: {{.Compare}}
		GitHub-User: {{.Pusher.Name}}
		{{- range $_, $c := .Commits -}}
			[info][title]{{- $c.Message -}}[/title]URL: {{ $c.URL -}}
			[hr]
			{{- if ne (len $c.Added) 0 -}}
				Added:[code]
				{{- range $_, $a := $c.Added -}}
					{{ $a }}
				{{- end -}}
				[/code]
			{{- end -}}
			{{- if ne (len $c.Modified) 0 -}}
				Modified:[code]
				{{- range $_, $m := $c.Modified -}}
					{{ $m }}
				{{- end -}}
				[/code]
			{{- end }}
			{{- if ne (len $c.Removed) 0 -}}
				Removed:[code]
				{{- range $_, $r := $c.Removed -}}
					{{ $r }}
				{{- end -}}
				[/code]
			{{- end -}}
			[/info] 
		{{- end -}}
		[/info]`

	tpl, err := template.New("Push").Parse(strings.Replace(templateText, "\t", "", -1))
	if err != nil {
		return "", err
	}

	msg := &bytes.Buffer{}

	if err := tpl.Execute(msg, pl); err != nil {
		return "", err
	}

	return msg.String(), nil
}

func (cl client) HandlePush(payload interface{}, c echo.Context) error {

	var res []byte

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

	if res, err = cl.cw.PostRoomMessage(cl.roomId, msg); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.String(http.StatusOK, string(res))
}
