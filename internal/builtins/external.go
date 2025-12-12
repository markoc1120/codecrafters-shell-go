package builtins

import (
	"os"
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
	cmd.Stderr = os.Stderr

	cmd.Run()
	return nil
}
