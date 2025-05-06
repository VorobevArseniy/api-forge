package spec

type Spec struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Method   string            `yaml:"method"`
	Path     string            `yaml:"path"`
	Request  map[string]string `yaml:"request"`
	Response map[string]string `yaml:"response"`
	Query    map[string]string `yaml:"query"`
}
