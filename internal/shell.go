package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"github.com/codecrafters-io/shell-starter-go/internal/register"
)

type MyShell struct {
	parser *parser.Parser
}

func CreateNewMyShell() *MyShell {
	reg := register.NewRegister(os.Stdout)
	register.RegisterBuiltins(reg)
	parser := parser.NewParser(reg)
	return &MyShell{parser: parser}
}

func (sh *MyShell) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())
		cmd, err := sh.parser.Parse(userInput)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if cmd == nil {
			continue
		}

		if err := cmd.Execute(); err != nil {
			fmt.Println(err)
		}

		cmd.CloseFile()
	}
}
