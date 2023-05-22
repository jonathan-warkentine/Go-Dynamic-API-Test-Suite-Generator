package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/jonathan-warkentine/Go-Dynamic-API-Test-Suite-Generator/models"
	"gopkg.in/yaml.v2"
)

func Unmarshal() ([]models.Group, error) {
	var tests []models.Group

	yamlFile, err := ioutil.ReadFile("tests.yml")
	if err != nil {
		return nil, fmt.Errorf("yamlFile.Get err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &tests)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal: %v", err)
	}

	return tests, nil
}
