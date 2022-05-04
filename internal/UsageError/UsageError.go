package UsageError

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
)

// HandleUsageError handles all errors
func HandleUsageError(ctx *cli.Context, err error, isSubcommand bool) error {
	color.Red(err.Error())
	os.Exit(0)
	return nil
}
