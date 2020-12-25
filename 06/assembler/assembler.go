package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var debugEnabled bool = false

func main() {
	// make sure we have the right number of arguments
	argError := checkArgs()

	if argError != "" {
		fmt.Println(argError)
		return
	}

	// initialize the input file
	absolutePath, input := initializeInput(os.Args[1])
	defer input.Close()

	// initialize the output file and writer
	output := initializeOutput(absolutePath)
	defer output.Flush()

	// first pass to identify labels
	scanForSymbols(input)

	// second pass to actually assemble
	input.Seek(0, 0)
	assemble(input, output)

	fmt.Printf("done! 😎\n")
}

func assemble(input *os.File, output *bufio.Writer) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		var outputLine string

		// remove comments
		line = ignoreComments(line)
		// ignore labels
		line = ignoreLabels(line)

		// if the line was a label, a comment, or whitespace, it will be a empty string
		if line != "" {
			if isAInstruction(line) {
				outputLine = handleAInstruction(line)
			} else {
				outputLine = handleCInstruction(line)
			}

			// write to output file
			fmt.Fprintln(output, outputLine)
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func debug(message string) {
	if debugEnabled {
		fmt.Println(message)
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

func ignoreLabels(line string) string {
	if isLabel(line) {
		line = ""
	}

	return line
}
