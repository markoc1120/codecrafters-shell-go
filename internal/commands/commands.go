package commands

import "io"

type Command interface {
	Execute() error
	SetStdout(io.Writer)
}

type BaseCommand struct {
	Stdout io.Writer
}

func (c *BaseCommand) SetStdout(output io.Writer) {
	c.Stdout = output
}
