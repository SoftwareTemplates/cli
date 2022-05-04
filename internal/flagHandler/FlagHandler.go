package flagHandler

import "github.com/urfave/cli/v2"

// InitGitFlag is used for setting if a new git repository should be initialized
var InitGitFlag = &cli.BoolFlag{
	Name:    "initGit",
	Aliases: []string{"ig", "git"},
	Usage:   "Use this flag if a git repository should be created on template init",
}

// TemplateFlag is used for selecting a template directly
var TemplateFlag = &cli.StringFlag{
	Name:    "template",
	Aliases: []string{"t"},
	Usage:   "The name of the template that should be used",
}

// ProjectNameFlag is used for selecting a project name directly from the command
var ProjectNameFlag = &cli.StringFlag{
	Name:    "projectName",
	Aliases: []string{"p"},
	Usage:   "The name of the project that should be used",
}

// HandleFlags returns all flags that are supported in the project
func HandleFlags() []cli.Flag {
	return []cli.Flag{
		InitGitFlag,
		TemplateFlag,
		ProjectNameFlag,
	}
}
