package builtins

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type EchoCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *EchoCommand) Execute() error {
	fmt.Fprintln(c.Stdout, strings.Join(c.Args, " "))
	return nil
}
