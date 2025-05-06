package generator

import (
	"api-generator/pkg/spec"
	"fmt"
	"os"
	"path/filepath"
)

type Generator struct {
	spec       *spec.Spec
	outputDir  string
	routerType string
}

func New(spec *spec.Spec, outputDir, routerType string) *Generator {
	return &Generator{
		spec:       spec,
		outputDir:  outputDir,
		routerType: routerType,
	}
}

func (g *Generator) Generate() error {
	if err := g.createDirs(); err != nil {
		return err
	}

	if err := g.GenerateAPIFiles(); err != nil {
		return err
	}

	if err := g.GenerateRoutes(); err != nil {
		return err
	}

	if err := g.GenerateInternalStubs(); err != nil {
		return err
	}

	return g.GenerateMain()
}

func (g *Generator) createDirs() error {
	dirs := []string{
		filepath.Join(g.outputDir, "api"),
		filepath.Join(g.outputDir, "routes"),
		filepath.Join(g.outputDir, "internal"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}
