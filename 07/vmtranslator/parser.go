package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Parser handles taking an input file, reading it,
// and parsing individual lines as separate commands.
type Parser struct {
	inputScanner                                  *bufio.Scanner
	hasMoreCommands                               bool
	currentLine, commandType, command, arg1, arg2 string
}

var commandTypes = map[string]string{
	"arithmetic": "ARITHMETIC",
}

func (p *Parser) advance() {
	p.clear()
	switch p.inputScanner.Scan() {
	case true:
		p.handleLine(p.inputScanner.Text())
	case false:
		p.hasMoreCommands = false
	}
}

func (p *Parser) handleLine(line string) {
	p.currentLine = ignoreComments(line)
	args := strings.Split(p.currentLine, " ")

	switch len(args) {
	case 1:
		command := args[0]

		if command == "" {
			return
		}

		p.commandType = getSingleCommandType(command)
		p.command, p.arg1, p.arg2 = command, "", ""
		return
	case 3:
		command := args[0]
		p.commandType = getCommnandWithArgType(command)
		p.command, p.arg1, p.arg2 = args[0], args[1], args[2]
		return
	default:
		panic(fmt.Sprintf("this command is not valid! %s", line))
	}
}

func (p *Parser) clear() {
	p.currentLine, p.commandType, p.command, p.arg1, p.arg2 = "", "", "", "", ""
}

func getSingleCommandType(command string) string {
	if command == "return" {
		return "C_RETURN"
	}

	return "C_ARITHMETIC"
}

func getCommnandWithArgType(command string) string {
	switch command {
	case "push":
		return "C_PUSH"
	case "pop":
		return "C_POP"
	default:
		panic(fmt.Sprintf("unknown command! %s", command))
	}
}

func ignoreComments(line string) string {
	hasComment := strings.Contains(line, "//")
	line = strings.TrimSpace(line)

	if hasComment {
		split := strings.Split(line, "//")

		if len(split) > 1 {
			// ignore trailing comments
			line = split[0]
		} else {
			// ignore whole line comments
			line = ""
		}
	}

	return strings.TrimSpace(line)
}
