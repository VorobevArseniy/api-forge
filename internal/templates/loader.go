package templates

import (
	"embed"
	"strings"
	"text/template"

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
		"ToTitle": caser.String,
		"ToLower": strings.ToLower,
		"ToCamel": toCamelCase,
		"toCamel": toCamelCase,
		"ToGoType": func(s string) string {
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

func toCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	for i := 1; i < len(words); i++ {
		words[i] = cases.Title(language.English).String(words[i])
	}

	return strings.Join(words, "")
}
