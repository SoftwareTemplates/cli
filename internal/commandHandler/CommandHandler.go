package commandHandler

import (
	"cli/internal/commands"
	"cli/internal/flagHandler"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// CommandHandler defines base functionality for
// a general purpose command handler
type CommandHandler interface {

	// HandleCommand builds all commands and returns
	// them as a pointer
	HandleCommand() []*cli.Command

	// NotImplemented handles the default response
	// if a command is not implemented
	NotImplemented() error
}

type InternalHandler struct{}

// HandleCommand builds all commands and returns
// them as a pointer
func (h InternalHandler) HandleCommand() []*cli.Command {
	return []*cli.Command{
		&cli.Command{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Creates a project base on a template",
			Action:  commands.InitCommand,
			Flags: []cli.Flag{
				flagHandler.InitGitFlag,
				flagHandler.TemplateFlag,
				flagHandler.ProjectNameFlag,
				flagHandler.CustomTemplateUrlFlag,
			},
		},
	}
}

// NotImplemented handles the default response
// if a command is not implemented
func (h InternalHandler) NotImplemented() error {
	color.Red("Error: This command does not exist")
	return nil
}
