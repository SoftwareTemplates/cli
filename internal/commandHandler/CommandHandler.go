package commandHandler

import (
	"cli/internal/commands"
	"cli/internal/flagHandler"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type CommandHandler interface {
	HandleCommand() []*cli.Command
	NotImplemented() error
}

type InternalHandler struct{}

func (h InternalHandler) HandleCommand() []*cli.Command {
	return []*cli.Command{
		&cli.Command{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Creates a project base on a template",
			Action:  commands.InitCommand,
			Flags:   []cli.Flag{flagHandler.InitGitFlag, flagHandler.TemplateFlag, flagHandler.ProjectNameFlag},
		},
	}
}

func (h InternalHandler) NotImplemented() error {
	color.Red("Error: This command does not exist")
	return nil
}
