package views

import (
	"embed"
	"html/template"
)

var (
	//go:embed *.html
	embedTemplate embed.FS

	funcMap       = template.FuncMap{}
	ViewsTemplate = template.Must(
		template.New("").
			Funcs(funcMap).
			ParseFS(embedTemplate, "*.html"))
)
