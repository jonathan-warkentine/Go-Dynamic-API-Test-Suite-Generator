package models

// type Header struct {

// }

type Test struct {
	name string `yaml: "name"`
	description string `yaml: "description"`
	url string `yaml: "url"`
	method string `yaml: "method"`
	headers []Header `yaml: "headers"`
}

type Group struct {
	group string `yaml: "group"`
	tests []Test `yaml: "tests"`
}
