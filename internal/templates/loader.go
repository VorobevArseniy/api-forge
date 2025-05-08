package templates

import (
	"bytes"
	"embed"
	"strings"
	"text/template"
	"unicode"

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
		"ToTitle":  caser.String,
		"ToLower":  strings.ToLower,
		"ToCamel":  toCamelCase,
		"ToPascal": toPascalCase,
		"ToSnake":  toSnakeCase,
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

func toPascalCase(s string) string {
	if s == "" {
		return s
	}

	pascal := toCamelCase(s)

	return string(unicode.ToUpper(rune(pascal[0]))) + pascal[1:]
}

func toSnakeCase(s string) string {
	var buf bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			// Add underscore before uppercase letters (except first character)
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else if r == '-' || r == ' ' {
			// Replace hyphens/spaces with underscores
			buf.WriteRune('_')
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
