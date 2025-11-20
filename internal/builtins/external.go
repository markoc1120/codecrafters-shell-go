package builtins

import (
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type ExternalCommand struct {
	commands.BaseCommand
	Cmd  string
	Args []string
}

func (c *ExternalCommand) Execute() error {
	cmd := exec.Command(c.Cmd, c.Args...)
	cmd.Stdout = c.Stdout
	return cmd.Run()
}
