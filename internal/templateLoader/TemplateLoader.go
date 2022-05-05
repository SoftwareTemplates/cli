package templateLoader

import (
	"cli/internal/formatter"
	"context"
	"errors"
	"github.com/google/go-github/github"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

// GetAllTemplates fetches all templates from the
// github api and leaves out non-template repos
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

// Checks if the template with the given name exists in the provided
// list of repositorys
func checkTemplateInArray(templates []*github.Repository, template string) bool {
	inArray := false
	for _, t := range templates {
		if *t.Name == template {
			inArray = true
		}
	}
	return inArray
}

// Gets a specific template from the given list by its name
func getTemplateByName(templates []*github.Repository, templateName string) *github.Repository {
	for _, t := range templates {
		if *t.Name == templateName {
			return t
		}
	}
	return nil
}

// GetTemplateName gets the name of the temlate
// If no cli argument is provided a promt will be opened
// to get the template name
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

// LoadTemplate handles all template and project specific data
// and finally saves the project and renames it
func LoadTemplate(templateName string, projectName string, customTemplate bool) error {

	templates := GetAllTemplates()
	if !checkTemplateInArray(templates, templateName) && !customTemplate {
		return errors.New("this template does not exist")
	}
	cmd := exec.Command("git", "clone", templateName)
	projectPath := formatter.FormatProjectName(templateName)
	if !customTemplate {
		templateRepo := getTemplateByName(templates, templateName)
		cmd = exec.Command("git", "clone", *templateRepo.HTMLURL)
		projectPath = *templateRepo.Name
	}
	err := cmd.Run()
	if err != nil {
		return err
	}
	err = os.RemoveAll("./" + projectPath + "/.git")
	if err != nil {
		return err
	}
	err = os.Rename("./"+projectPath, "./"+projectName)
	if err != nil {
		return err
	}
	return nil
}
