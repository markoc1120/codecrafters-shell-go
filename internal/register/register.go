package register

import (
	"fmt"
	"io"
	"maps"
	"os/exec"
	"slices"

	"github.com/codecrafters-io/shell-starter-go/internal/builtins"
	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

type CommandRegister struct {
	commands map[string]func(args []string, stdout io.WriteCloser) commands.Command
	stdout   io.WriteCloser
}

func NewRegister(stdout io.WriteCloser) *CommandRegister {
	return &CommandRegister{
		commands: make(map[string]func(args []string, stdout io.WriteCloser) commands.Command),
		stdout:   stdout,
	}
}

func RegisterBuiltins(register *CommandRegister) {
	register.Register("exit", func(args []string, stdout io.WriteCloser) commands.Command {
		return &builtins.ExitCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}}
	})
	register.Register("echo", func(args []string, stdout io.WriteCloser) commands.Command {
		return &builtins.EchoCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args}
	})
	register.Register("pwd", func(args []string, stdout io.WriteCloser) commands.Command {
		return &builtins.PwdCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}}
	})
	register.Register("cd", func(args []string, stdout io.WriteCloser) commands.Command {
		return &builtins.CdCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args}
	})

	builtinCmds := slices.Collect(maps.Keys(register.commands))
	builtinCmds = append(builtinCmds, "type")
	register.Register("type", func(args []string, stdout io.WriteCloser) commands.Command {
		return &builtins.TypeCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args, BuiltinCmds: builtinCmds}
	})
}

func (cr *CommandRegister) Register(cmd string, builtinFunc func(args []string, stdout io.WriteCloser) commands.Command) {
	cr.commands[cmd] = builtinFunc
}

func (cr *CommandRegister) Create(cmd string, args []string) (commands.Command, error) {
	if cmdFunc, found := cr.commands[cmd]; found {
		return cmdFunc(args, cr.stdout), nil
	} else if _, err := exec.LookPath(cmd); err == nil {
		return &builtins.ExternalCommand{Cmd: cmd, BaseCommand: commands.BaseCommand{Stdout: cr.stdout}, Args: args}, nil
	}
	return nil, fmt.Errorf("%s: command not found", cmd)
}
