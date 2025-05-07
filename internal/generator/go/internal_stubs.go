package generator

import (
	"api-generator/internal/templates"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (g *Generator) GenerateInternalStubs() error {
	tmpl, err := templates.Load("go/internal_stubs.go.tmpl")
	if err != nil {
		return err
	}

	for svcName, svc := range g.spec.Services {
		// Создаем папку для сервиса
		serviceDir := filepath.Join(g.outputDir, "internal", strings.ToLower(svcName))
		if err := os.MkdirAll(serviceDir, 0755); err != nil {
			return fmt.Errorf("failed to create service directory: %w", err)
		}

		// Готовим данные для шаблона
		data := struct {
			PackageName string
			Endpoints   []EndpointTemplate
			ModuleName  string
		}{
			PackageName: strings.ToLower(svcName),
			Endpoints:   convertEndpoints(svc.Endpoints),
			ModuleName:  g.moduleName,
		}

		// Генерируем и записываем файл
		content, err := templates.Render(tmpl, data)
		if err != nil {
			return fmt.Errorf("failed to render template: %w", err)
		}

		stubFile := filepath.Join(serviceDir, "service.go")
		if err := os.WriteFile(stubFile, content, 0644); err != nil {
			return fmt.Errorf("failed to write stub file: %w", err)
		}
	}
	return nil
}
