package generator

import (
	"api-generator/pkg/spec"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Generator struct {
	spec       *spec.Spec
	outputDir  string
	routerType string
	moduleName string
}

func New(spec *spec.Spec, outputDir, routerType, moduleName string) *Generator {
	return &Generator{
		spec:       spec,
		outputDir:  outputDir,
		routerType: routerType,
		moduleName: moduleName,
	}
}

func (g *Generator) Generate() error {
	if err := g.InitGoModule(); err != nil {
		return err
	}

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

	if err := g.GenerateMain(); err != nil {
		return err
	}

	return g.RunGoModTidy()
}

func (g *Generator) InitGoModule() error {
	if g.moduleName == "" {
		return fmt.Errorf("project module path is required")
	}

	cmd := exec.Command("go", "mod", "init", g.moduleName)
	cmd.Dir = g.outputDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod init failed: %s\n%w", string(output), err)
	}

	return nil
}

func (g *Generator) RunGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = g.outputDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod tidy failed: %s\n%w", string(output), err)
	}

	return nil
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
