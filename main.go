package main

import (
	"log"
	"os"
	"os/exec"
)

// Dynamic test suite can be generated and executed via `go run main.go` or `go test`
func main() {
	cmd := exec.Command("go", "test")
	cmd.Stdout = os.Stdout // connect the standard output of the cmd to the standard output of the program
	cmd.Stderr = os.Stderr // do the same for the standard error
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
