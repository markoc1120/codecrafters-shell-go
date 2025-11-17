package builtins

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type ExitCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *ExitCommand) Execute() error {
	os.Exit(0)
	return nil
}
