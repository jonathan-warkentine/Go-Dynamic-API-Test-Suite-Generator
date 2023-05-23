package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jonathan-warkentine/Go-Dynamic-API-Test-Suite-Generator/models"
	"github.com/jonathan-warkentine/Go-Dynamic-API-Test-Suite-Generator/utils"
)

// Function to execute an API call
func executeTest(t *testing.T, test models.Test, done chan bool) {
	defer func() { done <- true }() // Signal when this test is done

	client := &http.Client{} // Create an HTTP client

	req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(string(test.Body)))
	if err != nil {
		t.Fatalf("Failed to create %q request to %q: %v", test.Method, test.URL, err)
	}
	resp, err := client.Do(req) // Send the HTTP request
	if err != nil {
		t.Fatalf("Failed to send %q request to %q: %v", test.Method, test.URL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != test.Expect {
		t.Errorf("Expected status %d, got %d", test.Expect, resp.StatusCode)
	}
	fmt.Printf("PASS: %s (%s to %s)\n", test.Name, test.Method, test.URL)
}

// Function to execute a group of tests
func executeTestGroup(t *testing.T, group models.Group) {
	done := make(chan bool) // Channel to signal when a test is done

	// Execute each test in a separate goroutine
	for _, test := range group.Tests {
		go executeTest(t, test, done)
	}

	// Wait for all tests in this group to complete
	for range group.Tests {
		<-done
	}
}

func TestApiCalls(t *testing.T) {
	groups, err := utils.Unmarshal()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	for _, group := range groups {
		t.Run(group.Group, func(t *testing.T) {
			executeTestGroup(t, group)
		})
	}
}
