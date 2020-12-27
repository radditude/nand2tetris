package main

import (
	"bufio"
	"fmt"
)

// Codewriter handles translating VM commands to
// hack assembly code and writing the result
type Codewriter struct {
	outputWriter *bufio.Writer
	counter      int
}

func (c *Codewriter) writeArithmetic(line string, command string) {
	c.startCommand(line)

	switch command {
	case "add":
		c.setUpArithmetic1()
		c.writeOutputLine("M=D+M")
	case "and":
		c.setUpArithmetic1()
		c.writeOutputLine("M=D&M")
	case "eq":
		c.comparison("JEQ")
	case "gt":
		c.comparison("JLT")
	case "lt":
		c.comparison("JGT")
	case "or":
		c.setUpArithmetic1()
		c.writeOutputLine("M=D|M")
	case "neg":
		c.setUpArithmetic2()
		c.writeOutputLine("M=-M")
	case "not":
		c.setUpArithmetic2()
		c.writeOutputLine("M=!M")
	case "sub":
		c.setUpArithmetic1()
		c.writeOutputLine("M=M-D")
	default:
		panic("unknown command")
	}
}

func (c *Codewriter) writePushPop(line string, command string, arg1 string, arg2 string) {
	c.startCommand(line)
	if command == "push" {
		c.writeOutputLine(fmt.Sprintf("@%s", arg2))
		c.writeOutputLine("D=A\n@SP\nA=M\nM=D")
		c.incrementPointer()
	} else if command == "pop" {
		c.decrementPointer()
		c.writeOutputLine("D=M")
	} else {
		panic("command is not push or pop")
	}
}

// add a comment with the text of the original VM
// command, for debugging
func (c *Codewriter) startCommand(line string) {
	c.writeOutputLine(fmt.Sprintf("\n// %s", line))
}

func (c *Codewriter) incrementPointer() {
	c.writeOutputLine("@SP\nM=M+1")
}

func (c *Codewriter) decrementPointer() {
	c.writeOutputLine("@SP\nM=M-1")
}

// starts all arithmetic/logical commands except neg and not
func (c *Codewriter) setUpArithmetic1() {
	c.writeOutputLine("@SP")
	c.writeOutputLine("AM=M-1")
	c.writeOutputLine("D=M")
	c.writeOutputLine("@SP")
	c.writeOutputLine("A=M-1")
}

// starts neg and not
func (c *Codewriter) setUpArithmetic2() {
	c.writeOutputLine("@SP")
	c.writeOutputLine("A=M-1")
}

func (c *Codewriter) comparison(jumpCommand string) {
	isLabel := fmt.Sprintf("IS_%d", c.counter)
	isNotLabel := fmt.Sprintf("IS_NOT_%d", c.counter)
	c.setUpArithmetic1()
	c.writeOutputLine("D=D-M")
	c.writeOutputLine(fmt.Sprintf("@%s", isLabel))
	c.writeOutputLine(fmt.Sprintf("D;%s", jumpCommand))
	c.writeOutputLine("@SP")
	c.writeOutputLine("A=M-1")
	c.writeOutputLine("M=0")
	c.writeOutputLine(fmt.Sprintf("@%s", isNotLabel))
	c.writeOutputLine("0;JMP")
	c.writeOutputLine(fmt.Sprintf("(%s)", isLabel))
	c.writeOutputLine("@SP")
	c.writeOutputLine("A=M-1")
	c.writeOutputLine("M=-1")
	c.writeOutputLine(fmt.Sprintf("(%s)", isNotLabel))
	c.counter++
}

// write a line to the output file (takes care of \n)
func (c *Codewriter) writeOutputLine(line string) {
	fmt.Fprintln(c.outputWriter, line)
}
