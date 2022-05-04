package commands

import (
	"cli/internal/templateInit"
	"cli/internal/templateLoader"
	"github.com/urfave/cli/v2"
)

func InitCommand(ctx *cli.Context) error {

	name := templateLoader.GetTemplateName(ctx)
	err := templateLoader.LoadTemplate(name)
	if err != nil {
		return err
	}
	templateInit.InitGitRepoIfRequired(ctx, name)
	return nil
}
