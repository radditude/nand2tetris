package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var debugEnabled bool = false

func main() {
	filepath, err := checkArgs()

	if err != nil {
		fmt.Print(err)
		return
	}

	// get input file ready
	absoluteFilepath, inputFile, err := initializeInput(filepath)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer inputFile.Close()

	// get output file ready
	outputWriter, err := initializeOutput(absoluteFilepath)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer outputWriter.Flush()

	inputScanner := bufio.NewScanner(inputFile)
	parser := Parser{inputScanner, true, "", "", "", "", ""}
	codewriter := Codewriter{outputWriter}

	for parser.hasMoreCommands {
		parser.advance()

		if parser.command != "" {
			switch parser.commandType {
			case "C_ARITHMETIC":
				codewriter.writeArithmetic(parser.currentLine, parser.command)
			case "C_PUSH", "C_POP":
				codewriter.writePushPop(parser.currentLine,
					parser.command, parser.arg1, parser.arg2)
			}
		}
	}

	fmt.Printf("\nWe did it! 🤘🤘🤘\n")
}

func checkArgs() (string, error) {
	argsLength := len(os.Args)

	if argsLength > 2 {
		return "", errors.New("please provide only one file name")
	}

	if argsLength < 2 {
		return "", errors.New("please provide a file name")
	}

	return os.Args[1], nil
}

func initializeInput(inputPath string) (string, *os.File, error) {
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return "", nil, err
	}

	return absPath, file, nil
}

func initializeOutput(inputFile string) (*bufio.Writer, error) {
	path, fileName := filepath.Split(inputFile)
	basename := strings.Split(fileName, ".")[0]
	outputFileName := fmt.Sprintf("%s%s.asm", path, basename)

	output, err := os.Create(outputFileName)
	if err != nil {
		return nil, err
	}

	return bufio.NewWriter(output), nil
}

func debug(message string) {
	if debugEnabled {
		fmt.Println(message)
	}
}
