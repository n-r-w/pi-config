package formexample

import (
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templateFiles embed.FS

func parseTemplates() (*template.Template, error) {
	return template.New("contacts").ParseFS(templateFiles, "templates/*.html")
}
