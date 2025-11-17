package parser

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
	"github.com/codecrafters-io/shell-starter-go/internal/register"
)

type Parser struct {
	commandRegister *register.CommandRegister
}

func NewParser(reg *register.CommandRegister) *Parser {
	return &Parser{commandRegister: reg}
}

func (p *Parser) Parse(input string) (commands.Command, error) {
	inputs := strings.Split(input, " ")
	cmd, args := inputs[0], inputs[1:]

	cmdFunc, err := p.commandRegister.Create(cmd, args)
	if err != nil {
		return nil, err
	}
	return cmdFunc, nil
}
