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

func parseArguments(inputs []string) []string {
	if len(inputs) == 1 {
		return nil
	}

	var result []string

	var argument strings.Builder
	isSingleQuote := false
	inputString := inputs[1]
	inputStringLen := len(inputString)

	for idx, char := range inputString {
		// tracking single quotes, saving parsed argument
		if char == '\'' {
			if isSingleQuote && idx+1 <= inputStringLen-1 {
				nextRune := inputString[idx+1]
				if nextRune == ' ' {
					result = append(result, argument.String())
					argument.Reset()
				}
			}
			if idx == inputStringLen-1 {
				result = append(result, argument.String())
				argument.Reset()
			}
			isSingleQuote = !isSingleQuote
			continue
		}

		if isSingleQuote {
			argument.WriteRune(char)
			continue
		}

		if char != ' ' {
			argument.WriteRune(char)
			if idx+1 <= inputStringLen-1 {
				nextRune := inputString[idx+1]
				if nextRune == ' ' {
					result = append(result, argument.String())
					argument.Reset()
				}
			}
		}

		if idx == inputStringLen-1 {
			result = append(result, argument.String())
			argument.Reset()
		}
	}
	return result
}

func (p *Parser) Parse(input string) (commands.Command, error) {
	inputs := strings.SplitN(input, " ", 2)

	cmd := inputs[0]
	args := parseArguments(inputs)

	cmdFunc, err := p.commandRegister.Create(cmd, args)
	if err != nil {
		return nil, err
	}
	return cmdFunc, nil
}
