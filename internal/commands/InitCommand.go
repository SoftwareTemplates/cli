package commands

import (
	"cli/internal/templateInit"
	"cli/internal/templateLoader"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func getProjectName(ctx *cli.Context) string {

	if ctx.String("projectName") != "" {
		return ctx.String("projectName")
	}
	prompt := promptui.Prompt{
		Label: "Projectname",
		Validate: func(s string) error {
			return nil
		},
	}
	projectName, err := prompt.Run()
	if err != nil {
		panic(err.Error())
	}
	return projectName
}

func InitCommand(ctx *cli.Context) error {
	projectName := getProjectName(ctx)
	name := templateLoader.GetTemplateName(ctx)
	err := templateLoader.LoadTemplate(name, projectName)
	if err != nil {
		return err
	}
	templateInit.InitGitRepoIfRequired(ctx, projectName)
	return nil
}
