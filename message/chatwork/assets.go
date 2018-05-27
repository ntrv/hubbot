package chatwork

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsdaa4d3d0a9722b3ae57f515a4935214b03826800 = "[info][title]Push event at {{.Repository.Name}}[/title]URL: {{.Repository.URL}}\n  Ref: {{.Ref}}\n  CompareURL: {{.Compare}}\n  GitHub-User: {{.Pusher.Name}}\n  {{- range $_, $c := .Commits -}}\n    [info][title]{{- $c.Message -}}[/title]URL: {{ $c.URL -}}\n    [hr]\n    {{- if ne (len $c.Added) 0 -}}\n      Added:[code]\n      {{- range $_, $a := $c.Added -}}\n        {{ $a }}\n      {{- end -}}\n      [/code]\n    {{- end -}}\n    {{- if ne (len $c.Modified) 0 -}}\n      Modified:[code]\n      {{- range $_, $m := $c.Modified -}}\n        {{ $m }}\n      {{- end -}}\n      [/code]\n    {{- end }}\n    {{- if ne (len $c.Removed) 0 -}}\n      Removed:[code]\n      {{- range $_, $r := $c.Removed -}}\n        {{ $r }}\n      {{- end -}}\n      [/code]\n    {{- end -}}\n    [/info] \n  {{- end -}}\n[/info]\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"push.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1527408919, 1527408919000000000),
		Data:     nil,
	}, "/push.tmpl": &assets.File{
		Path:     "/push.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1527408919, 1527408919000000000),
		Data:     []byte(_Assetsdaa4d3d0a9722b3ae57f515a4935214b03826800),
	}}, "")
