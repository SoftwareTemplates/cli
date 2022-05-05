#! /bin/sh

GOOS=linux GOARCH=arm64 go build -o "softwareTemplates-macOS-m1" cmd/main.go
GOOS=linux go build -o "softwareTemplates-linux-macOS" cmd/main.go
GOOS=windows go build -o "softwareTemplates-windows.exe" cmd/main.go