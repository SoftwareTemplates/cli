package flagHandler

import "github.com/urfave/cli/v2"

func HandleFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "initGit",
			Aliases: []string{"ig", "git"},
			Usage:   "Use this flag if a git repository should be created on template init",
		},
		&cli.StringFlag{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "The name of the template that should be used",
		},
	}
}
