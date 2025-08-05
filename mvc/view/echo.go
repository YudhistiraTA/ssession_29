package view

import (
	_ "embed"
	"html/template"
)

//go:embed template/echo.html
var echoHTML string

func EchoTemplate() *template.Template {
	return template.Must(template.New("echo").Parse(echoHTML))
}
