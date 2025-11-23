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

func (c *CdCommand) changeHome() error {
	val, found := os.LookupEnv("HOME")
	if found {
		os.Chdir(val)
		return nil
	}
	return fmt.Errorf("no HOME is defined")
}

func (c *CdCommand) Execute() error {
	if len(c.Args) == 0 {
		return c.changeHome()
	}

	requestedDir := c.Args[0]
	if requestedDir == "~" {
		c.changeHome()
	} else if err := os.Chdir(requestedDir); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", requestedDir)
	}
	return nil
}
