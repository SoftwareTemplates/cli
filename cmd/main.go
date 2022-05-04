package main

import (
	"cli/internal/commandHandler"
	"cli/internal/flagHandler"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	handler := commandHandler.InternalHandler{}
	app := &cli.App{
		Name:        "SoftareTemplates",
		Description: "A simple cli tool for making initializing projects faster",
		Commands:    handler.HandleCommand(),
		Flags:       flagHandler.HandleFlags(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
