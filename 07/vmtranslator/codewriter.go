package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
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
	switch command {
	case "push":
		c.push(arg1, arg2)
	case "pop":
		c.pop(arg1, arg2)
	default:
		panic("unknown command")
	}
}

// add a comment with the text of the original VM
// command, for debugging
func (c *Codewriter) startCommand(line string) {
	c.writeOutputLine(fmt.Sprintf("\n// %s", line))
}

func (c *Codewriter) push(arg1 string, arg2 string) {
	isConstant := arg1 == "constant"
	isDynamicSegment := isDynamicSegment(arg1)

	if !isDynamicSegment {
		c.writeOutputLine(getSegment(arg1, arg2))
		c.writeOutputLine("D=M")
		c.writeOutputLine("@SP")
		c.writeOutputLine("A=M")
	} else {
		c.writeOutputLine(fmt.Sprintf("@%s", arg2))
		c.writeOutputLine("D=A")
		c.writeOutputLine(getSegment(arg1, arg2))
		c.writeOutputLine("A=M")
	}

	if !isConstant && isDynamicSegment {
		c.writeOutputLine("AD=D+A")
		c.writeOutputLine("D=M")
		c.writeOutputLine("@SP")
		c.writeOutputLine("A=M")
	}

	c.writeOutputLine("M=D")
	c.writeOutputLine("@SP")
	c.writeOutputLine("M=M+1")
}

func (c *Codewriter) pop(arg1 string, arg2 string) {
	isDynamicSegment := isDynamicSegment(arg1)
	c.writeOutputLine("@SP")
	c.writeOutputLine("M=M-1")

	if isDynamicSegment {
		c.writeOutputLine(fmt.Sprintf("@%s", arg2))
		c.writeOutputLine("D=A")
		c.writeOutputLine(getSegment(arg1, arg2))
		c.writeOutputLine("A=M")
		c.writeOutputLine("D=D+A")
		c.writeOutputLine("@R13")
		c.writeOutputLine("M=D")
		c.writeOutputLine("@SP")
	}

	c.writeOutputLine("A=M")
	c.writeOutputLine("D=M")

	if isDynamicSegment {
		c.writeOutputLine("@R13")
		c.writeOutputLine("A=M")
	} else {
		c.writeOutputLine(getSegment(arg1, arg2))
	}

	c.writeOutputLine("M=D")
}

// TODO: there's probably a better way to do this
func isDynamicSegment(segment string) bool {
	re := regexp.MustCompile(segment)
	return !re.MatchString("staticpointertemp")
}

func getSegment(segment string, position string) string {
	if isDynamicSegment(segment) {
		segments := map[string]string{
			"constant": "@SP",
			"local":    "@LCL",
			"argument": "@ARG",
			"this":     "@THIS",
			"that":     "@THAT",
		}
		return segments[segment]
	}

	startingPoints := map[string]int64{
		"pointer": 3,
		"temp":    5,
		"static":  16,
	}

	startingPoint := startingPoints[segment]
	positionInt, _ := strconv.ParseInt(position, 10, 64)
	return fmt.Sprintf("@%d", startingPoint+positionInt)
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
