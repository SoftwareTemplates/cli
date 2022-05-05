package formatter

import "strings"

// FormatProjectName checks if the project name contains slashes.
// All slashes are removed and the words after the last slash
// will be used as project name
func FormatProjectName(projectName string) string {
	if strings.Contains(projectName, "/") {
		splitted := strings.Split(projectName, "/")
		return splitted[len(splitted)-1]
	}
	return projectName
}
