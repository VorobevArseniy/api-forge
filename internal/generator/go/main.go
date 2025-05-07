package generator

import (
	"api-generator/internal/templates"
	"api-generator/pkg/spec"
	"os"
	"path/filepath"
)

func (g *Generator) GenerateMain() error {
	tmpl, err := templates.Load("go/main.go.tmpl")
	if err != nil {
		return err
	}

	data := struct {
		RouterType string
		Services   map[string]spec.Service
		ModuleName string
	}{
		g.routerType,
		g.spec.Services,
		g.moduleName,
	}

	content, err := templates.Render(tmpl, data)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(g.outputDir, "main.go"), content, 0644)
}
