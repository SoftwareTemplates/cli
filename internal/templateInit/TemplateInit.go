package templateInit

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"os/exec"
)

// Checks if a git repository should be created
// If there is no cli ingut given a promt will be opened
func getShouldCreateGit(ctx *cli.Context) bool {
	if ctx.Bool("initGit") {
		return true
	}
	prompt := promptui.Select{
		Label: "Should a git repository be created?",
		Items: []string{"yes", "no"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		color.Red("Prompt input failed")
	}
	if result == "yes" {
		return true
	}
	return false
}

// InitGitRepoIfRequired executes the init git action and handles if the
// repository should be initialized
func InitGitRepoIfRequired(ctx *cli.Context, projectName string) {

	createGit := getShouldCreateGit(ctx)
	if createGit {
		fmt.Println("create")
		cmd := exec.Command("git", "init", projectName)
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	}
}
