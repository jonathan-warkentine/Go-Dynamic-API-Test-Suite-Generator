package main

import (
	"fmt"
	"log"

	"github.com/jonathan-warkentine/Go-Dynamic-API-Test-Suite-Generator/utils"
)

func main() {
	groups, err := utils.Unmarshal()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, group := range groups {
		fmt.Printf("Group: %+v\n", group)
	}
}
