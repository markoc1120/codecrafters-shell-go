package builtins

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type CdCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *CdCommand) Execute() error {
	if len(c.Args) == 0 {
		return fmt.Errorf("error for now")
	}
	requestedDir := c.Args[0]
	if err := os.Chdir(requestedDir); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", requestedDir)
	}
	return nil
}
