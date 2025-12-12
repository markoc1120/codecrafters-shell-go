package parser

import (
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
	"github.com/codecrafters-io/shell-starter-go/internal/register"
)

var (
	stdoutRedirects = []string{">", "1>"}
)

func parseArguments(inputString string) ([]string, error) {
	t := newTokenizer(strings.NewReader(inputString))
	arguments := make([]string, 0)

	for {
		arg, err := t.Next()
		if err == io.EOF {
			if arg != "" {
				arguments = append(arguments, arg)
			}
			return arguments, nil
		}
		if err != nil {
			return nil, err
		}
		if arg != "" {
			arguments = append(arguments, arg)
		}
	}
}

func parseRedirects(args []string) ([]string, string, error) {
	var redirectPath string
	for _, stdRedirect := range stdoutRedirects {
		idx := slices.Index(args, stdRedirect)
		// no redirect found
		if idx == -1 {
			continue
		}

		// otherwise update args
		redirectPath = args[idx+1]
		args = args[:idx]
		break
	}
	return args, redirectPath, nil
}

type Parser struct {
	commandRegister *register.CommandRegister
}

func NewParser(reg *register.CommandRegister) *Parser {
	return &Parser{commandRegister: reg}
}

func (p *Parser) Parse(input string) (commands.Command, error) {
	parsedArgs, err := parseArguments(input)
	if err != nil {
		return nil, err
	}

	var cmd string
	var args []string
	if len(parsedArgs) == 1 {
		cmd = parsedArgs[0]
	} else if len(parsedArgs) > 1 {
		cmd, args = parsedArgs[0], parsedArgs[1:]
	}

	args, redirectPath, err := parseRedirects(args)
	cmdFunc, err := p.commandRegister.Create(cmd, args)
	if err != nil {
		return nil, err
	}

	if redirectPath != "" {
		if err := os.MkdirAll(filepath.Dir(redirectPath), 0770); err != nil {
			return nil, err
		}

		outfile, err := os.Create(redirectPath)
		if err != nil {
			return nil, err
		}
		cmdFunc.SetStdout(outfile)
	}
	return cmdFunc, nil
}
