package generator

import (
	"api-generator/internal/templates"
	"api-generator/pkg/spec"
	"fmt"
	"os"
	"path/filepath"
)

func (g *Generator) GenerateAPIFiles() error {
	if err := g.generateModels(); err != nil {
		return fmt.Errorf("models generation failed: %w", err)
	}

	if err := g.generateInterfaces(); err != nil {
		return fmt.Errorf("interfaces generation failed: %w", err)
	}

	return g.generateHandlers()
}

func (g *Generator) generateModels() error {
	tmpl, err := templates.Load("go/models.go.tmpl")
	if err != nil {
		return nil
	}

	data := struct{ Services map[string]spec.Service }{g.spec.Services}

	content, err := templates.Render(tmpl, data)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(g.outputDir, "api/models.go"), content, 0644)
}

func (g *Generator) generateInterfaces() error {
	tmpl, err := templates.Load("go/interfaces.go.tmpl")
	if err != nil {
		return err
	}

	data := struct{ Services map[string]spec.Service }{g.spec.Services}

	content, err := templates.Render(tmpl, data)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(g.outputDir, "api/interfaces.go"), content, 0644)
}

func (g *Generator) generateHandlers() error {
	tmpl, err := templates.Load("go/handlers.go.tmpl")
	if err != nil {
		return err
	}

	data := struct{ Services map[string]spec.Service }{g.spec.Services}

	content, err := templates.Render(tmpl, data)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(g.outputDir, "api/handlers.go"), content, 0644)
}
