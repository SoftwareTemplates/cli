package commands

import (
	"cli/internal/ScriptExecutor"
	"cli/internal/formatter"
	"cli/internal/templateInit"
	"cli/internal/templateLoader"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

// Gets the project name from the input
// If no cli argument is given it will use a promt
// for selecting the project name
func getProjectName(ctx *cli.Context) string {

	// Check if cli argument exists
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

// InitCommand executes all methods that are
// required for initializing a project from template
func InitCommand(ctx *cli.Context) error {
	projectName := getProjectName(ctx)
	name := templateLoader.GetTemplateName(ctx)

	// Clones the project and deletes the old git repo
	err := templateLoader.LoadTemplate(name, formatter.FormatProjectName(projectName))
	if err != nil {
		return err
	}

	ScriptExecutor.ExecuteScript(projectName)

	// Creates a new git repo of it is required
	templateInit.InitGitRepoIfRequired(ctx, formatter.FormatProjectName(projectName))
	return nil
}
