package ScriptExecutor

import (
	"cli/internal/formatter"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
)

// ExecuteScript checks if there are init scripts given in the template
// if the scripts are given the project will be initialized by them
func ExecuteScript(projectName string) {
	bashScript := true
	shellScript := true
	formattedProjectName := formatter.FormatProjectName(projectName)
	if _, err := os.Stat("./" + formattedProjectName + "/init.bat"); errors.Is(err, os.ErrNotExist) {
		bashScript = false
	}
	if _, err := os.Stat("./" + formattedProjectName + "/init.sh"); errors.Is(err, os.ErrNotExist) {
		shellScript = false
	}
	if runtime.GOOS == "windows" && bashScript {
		cmd := exec.Command(
			"./"+formattedProjectName+"/init.bat",
			formattedProjectName,
			projectName)
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	} else if (runtime.GOOS == "linux" || runtime.GOOS == "darwin") && shellScript {
		cmd := exec.Command(
			"sh",
			"./"+formattedProjectName+"/init.sh",
			formattedProjectName,
			projectName)
		fmt.Println(cmd.String())
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	} else {
		return
	}
	if shellScript {
		err := os.Remove("./" + formattedProjectName + "/init.sh")
		if err != nil {
			panic(err.Error())
		}
	}
	if bashScript {
		err := os.Remove("./" + formattedProjectName + "/init.bat")
		if err != nil {
			panic(err.Error())
		}
	}
	color.Green("Initialized project infrastructure")
}
