package config

import (
	"api-generator/pkg/spec"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseSpec(path string) (*spec.Spec, error) {
	data, _ := os.ReadFile(path)
	spec := &spec.Spec{}

	err := yaml.Unmarshal(data, spec)
	if err != nil {
		return nil, err
	}

	return spec, nil
}
