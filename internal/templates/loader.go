package templates

import (
	"embed"
	"html/template"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed go/*.tmpl ts/*.tmpl
var tmplFS embed.FS // Встроенные шаблоны

func Load(name string) (*template.Template, error) {
	data, err := tmplFS.ReadFile(name)
	if err != nil {
		return nil, err
	}

	caser := cases.Title(language.English)

	tmpl := template.New(name).Funcs(template.FuncMap{
		"title":   caser.String,
		"toLower": strings.ToLower,
		"goType": func(s string) string {
			switch s {
			case "string":
				return "string"
			case "number":
				return "int"
			default:
				return "interface{}"
			}
		},
	})

	return tmpl.Parse(string(data))
}
