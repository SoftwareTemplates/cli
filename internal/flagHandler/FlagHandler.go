package flagHandler

import "github.com/urfave/cli/v2"

var InitGitFlag = &cli.BoolFlag{
	Name:    "initGit",
	Aliases: []string{"ig", "git"},
	Usage:   "Use this flag if a git repository should be created on template init",
}

var TemplateFlag = &cli.StringFlag{
	Name:    "template",
	Aliases: []string{"t"},
	Usage:   "The name of the template that should be used",
}

var ProjectNameFlag = &cli.StringFlag{
	Name:    "projectName",
	Aliases: []string{"p"},
	Usage:   "The name of the project that should be used",
}

func HandleFlags() []cli.Flag {
	return []cli.Flag{
		InitGitFlag,
		TemplateFlag,
		ProjectNameFlag,
	}
}
