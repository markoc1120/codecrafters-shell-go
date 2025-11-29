package parser

import (
	"io"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
	"github.com/codecrafters-io/shell-starter-go/internal/register"
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

	cmdFunc, err := p.commandRegister.Create(cmd, args)
	if err != nil {
		return nil, err
	}
	return cmdFunc, nil
}
