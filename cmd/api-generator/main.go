package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"api-generator/internal/config"
	generator "api-generator/internal/generator/go"
)

func main() {
	// Парсинг флагов командной строки
	inputFile := flag.String("input", "", "Path to YAML specification file")
	outputDir := flag.String("output", "./generated", "Output directory for generated files")
	routerType := flag.String("router", "chi", "Router type (chi or std)")
	projectModule := flag.String("module", "", "Go module path (e.g. 'github.com/username/project')")
	flag.Parse()

	// Валидация обязательных параметров
	if *inputFile == "" {
		log.Fatal("Error: input YAML file path is required")
	}
	if *projectModule == "" {
		log.Fatal("Error: project module path is required")
	}

	// Создание выходной директории
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Парсинг YAML-спецификации
	spec, err := config.ParseSpec(*inputFile)
	if err != nil {
		log.Fatalf("Failed to parse YAML spec: %v", err)
	}

	// Конфигурация генератора

	// Инициализация генератора
	gen := generator.New(spec, *outputDir, *routerType)

	// Выполнение всех генераций
	if err := gen.Generate(); err != nil {
		log.Fatalf("Generation failed: %v", err)
	}

	// Вывод информации о результате
	fmt.Printf(`
Successfully generated API code!
--------------------------------
Output directory: %s
Generated files:
  - API models:       %s/api/models.go
  - API interfaces:   %s/api/interfaces.go
  - API handlers:     %s/api/handlers.go
  - Router (%s):    %s/routes_%s.go
  - Main entrypoint:  %s/main.go
  - Service stubs:    %s/internal/{service}/

You can now implement your business logic in the generated stubs.
`,
		*outputDir,
		*outputDir, *outputDir, *outputDir,
		*routerType, *outputDir, *routerType,
		*outputDir,
		*outputDir,
	)
}
