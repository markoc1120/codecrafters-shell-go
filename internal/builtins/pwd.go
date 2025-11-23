package builtins

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type PwdCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *PwdCommand) Execute() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(c.Stdout, dir)
	return nil
}
