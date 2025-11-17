package builtins

import (
	"fmt"
	"slices"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type TypeCommand struct {
	commands.BaseCommand
	Args        []string
	BuiltinCmds []string
}

func (c *TypeCommand) Execute() error {
	if len(c.Args) == 0 {
		return fmt.Errorf("")
	}

	queryCmd := c.Args[0]
	found := slices.Contains(c.BuiltinCmds, queryCmd)

	var msg string
	if found {
		msg = fmt.Sprintf("%s is a shell builtin", queryCmd)
	} else {
		msg = fmt.Sprintf("%s: not found", queryCmd)
	}

	fmt.Fprintln(c.Stdout, msg)
	return nil
}
