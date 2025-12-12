package commands

import (
	"io"
	"os"
)

type Command interface {
	Execute() error
	SetStdout(io.WriteCloser)
	CloseFile()
}

type BaseCommand struct {
	Stdout io.WriteCloser
}

func (c *BaseCommand) SetStdout(output io.WriteCloser) {
	c.Stdout = output
}

func (c *BaseCommand) CloseFile() {
	if c.Stdout != os.Stdout {
		c.Stdout.Close()
	}
}
