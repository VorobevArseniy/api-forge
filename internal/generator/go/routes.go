package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"api-generator/internal/templates"
	"api-generator/pkg/spec"
)

func (g *Generator) GenerateRoutes() error {
	tmpl, err := templates.Load(fmt.Sprintf("go/routes_%s.go.tmpl", g.routerType))
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	data := struct {
		Services   map[string]spec.Service
		RouterType string
		ModuleName string
	}{
		Services:   g.spec.Services,
		RouterType: g.routerType,
		ModuleName: g.moduleName,
	}

	content, err := templates.Render(tmpl, data)
	if err != nil {
		return fmt.Errorf("failed to render routes: %w", err)
	}

	filename := filepath.Join(g.outputDir, fmt.Sprintf("routes/routes_%s.go", g.routerType))
	return os.WriteFile(filename, content, 0644)
}
