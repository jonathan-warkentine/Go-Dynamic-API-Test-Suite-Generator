# Dynamic API Test Suite Generator

## Description
Use `test.yml` to dynamically generate a suite of API tests that execute concurrently/in parallel for lightning execution. 

## Features
- Leverages Golang's goroutines for concurrent test/thread execution
- Test specifications are input via a simple yaml file, allowing for maximum flexibility: both in being able to adjust on the fly, and tracking changes via git

## Usage
Ensure that you have Go installed on your local machine. Fill out the `tests.yml` with your test specifications. Just run `go test` to execute your API tests!

This lightweight, lightspeed app is perfect for CI/CD pipeline deployments -- use it to test your endpoints after a successful deployment!