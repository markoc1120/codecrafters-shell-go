package register

import (
	"fmt"
	"io"
	"maps"
	"slices"

	"github.com/codecrafters-io/shell-starter-go/internal/builtins"
	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type CommandRegister struct {
	commands map[string]func(args []string, stdout io.Writer) commands.Command
	stdout   io.Writer
}

func NewRegister(stdout io.Writer) *CommandRegister {
	return &CommandRegister{
		commands: make(map[string]func(args []string, stdout io.Writer) commands.Command),
		stdout:   stdout,
	}
}

func RegisterBuiltins(register *CommandRegister) {
	register.Register("exit", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.ExitCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}}
	})
	register.Register("echo", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.EchoCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args}
	})

	builtinCmds := slices.Collect(maps.Keys(register.commands))
	builtinCmds = append(builtinCmds, "type")
	register.Register("type", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.TypeCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args, BuiltinCmds: builtinCmds}
	})
}

func (cr *CommandRegister) Register(cmd string, builtinFunc func(args []string, stdout io.Writer) commands.Command) {
	cr.commands[cmd] = builtinFunc
}

func (cr *CommandRegister) Create(cmd string, args []string) (commands.Command, error) {
	if cmdFunc, found := cr.commands[cmd]; found {
		return cmdFunc(args, cr.stdout), nil
	}
	return nil, fmt.Errorf("%s command not found", cmd)
}
