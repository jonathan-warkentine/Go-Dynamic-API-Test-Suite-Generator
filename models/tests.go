package models

type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Test struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	URL         string   `yaml:"url"`
	Method      string   `yaml:"method"`
	Headers     map[string]string `yaml:"headers"`
	Body		map[string]string `yaml:"body"`
}

type Group struct {
	Group string `yaml:"group"`
	Tests []Test `yaml:"tests"`
}