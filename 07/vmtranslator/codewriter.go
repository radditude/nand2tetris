package main

import (
	"bufio"
	"fmt"
)

// Codewriter handles generating assembly for a
// given command, then writing it to an output file
type Codewriter struct {
	outputWriter *bufio.Writer
}

func (c *Codewriter) writeArithmetic(line string, command string) {
	c.startCommand(line)
	c.writeOutputLine("@0\nM=M-1\n@0\nA=M\nD=M\nM=0\n@0\nM=M-1\nA=M\nM=M+D\n@0\nM=M+1")
}

func (c *Codewriter) writePushPop(line string, command string, arg1 string, arg2 string) {
	c.startCommand(line)
	if command == "push" {
		c.writeOutputLine(fmt.Sprintf("@%s", arg2))
		c.writeOutputLine("D=A\n@0\nA=M\nM=D\n@0\nM=M+1\n")
	} else if command == "pop" {
		// implement pop
	} else {
		panic("command is not push or pop")
	}
}

func (c *Codewriter) startCommand(line string) {
	c.writeOutputLine(fmt.Sprintf("// %s", line))
}

func (c *Codewriter) writeOutputLine(line string) {
	fmt.Fprintln(c.outputWriter, line)
}
