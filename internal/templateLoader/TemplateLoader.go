package templateLoader

import (
	"context"
	"errors"
	"github.com/google/go-github/github"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func GetAllTemplates() []*github.Repository {
	client := github.NewClient(nil)
	options := &github.RepositoryListOptions{}
	repos, _, _ := client.Repositories.List(context.Background(), "SoftwareTemplates", options)
	var templates []*github.Repository
	for _, repo := range repos {
		if *repo.Name != "cli" {
			templates = append(templates, repo)
		}
	}
	return templates
}

func checkTemplateInArray(templates []*github.Repository, template string) bool {
	inArray := false
	for _, t := range templates {
		if *t.Name == template {
			inArray = true
		}
	}
	return inArray
}
func getTemplateByName(templates []*github.Repository, templateName string) *github.Repository {
	for _, t := range templates {
		if *t.Name == templateName {
			return t
		}
	}
	return nil
}

func GetTemplateName(ctx *cli.Context) string {
	if ctx.String("template") != "" {
		return ctx.String("template")
	}
	templates := GetAllTemplates()
	var options []string
	for _, option := range templates {
		options = append(options, *option.Name)
	}
	promt := promptui.Select{
		Label: "Select a template",
		Items: options,
	}
	_, name, err := promt.Run()
	if err != nil {
		panic(err.Error())
	}
	return name
}

func LoadTemplate(templateName string, projectName string) error {

	templates := GetAllTemplates()
	if !checkTemplateInArray(templates, templateName) {
		return errors.New("this template does not exist")
	}
	templateRepo := getTemplateByName(templates, templateName)

	cmd := exec.Command("git", "clone", *templateRepo.HTMLURL)
	err := cmd.Run()
	if err != nil {
		return err
	}
	err = os.RemoveAll("./" + *templateRepo.Name + "/.git")
	if err != nil {
		return err
	}
	err = os.Rename("./"+*templateRepo.Name, "./"+projectName)
	if err != nil {
		return err
	}
	return nil
}
