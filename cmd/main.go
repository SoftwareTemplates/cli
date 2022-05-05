package main

import (
	"cli/internal/UsageError"
	"cli/internal/commandHandler"
	"cli/internal/flagHandler"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Inits the main console application
func main() {

	// Main command handler
	handler := commandHandler.InternalHandler{}

	app := &cli.App{
		Name:         "SoftareTemplates",
		Description:  "A simple cli tool for making initializing projects faster",
		Commands:     handler.HandleCommand(),
		Flags:        flagHandler.HandleFlags(),
		OnUsageError: UsageError.HandleUsageError,
		Copyright:    "2022 Mathis Burger",
		HelpName:     "Setup your projects faster",
		Version:      "1.3.0",
	}

	// Starts the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
