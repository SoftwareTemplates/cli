package commands

import (
	"cli/internal/templateLoader"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func TemplatesCommand(ctx *cli.Context) error {
	templates := templateLoader.GetAllTemplates()
	for _, template := range templates {
		green := color.New(color.FgGreen).Add(color.Bold)
		green.Print(*template.Name)
		fmt.Print(" - ")
		cyan := color.New(color.FgHiBlue)
		cyan.Println(*template.Description)
	}
	return nil
}
