package view

import (
	_ "embed"
	"html/template"
)

//go:embed template/home.html
var homeHTML string

func HomeTemplate() *template.Template {
	return template.Must(template.New("home").Parse(homeHTML))
}
