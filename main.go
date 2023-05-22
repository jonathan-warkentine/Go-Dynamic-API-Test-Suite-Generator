package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jonathan-warkentine/Go-Dynamic-API-Test-Suite-Generator/models"
	"gopkg.in/yaml.v2"
)

var tests []models.Group

func main() {
	yamlFile, err := ioutil.ReadFile("tests.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("Value: %#v\n", config)
}
