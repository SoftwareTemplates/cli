#! /bin/sh

GOOS=linux GOARCH=arm64 go build -o "softwareTemplates (macOS)" cmd/main.go
GOOS=linux go build -o "softwareTemplates (linux)" cmd/main.go
GOOS=windows go build -o "softwareTemplates (windows)" cmd/main.go