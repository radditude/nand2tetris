package assembler

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func checkArgs() string {
	argsLength := len(os.Args)
	var message string

	if argsLength > 2 {
		message = fmt.Sprintf("ERROR: Please provide only one file name.")
	}

	if argsLength < 2 {
		message = fmt.Sprintf("ERROR: Please provide a file name.")
	}

	return message
}

func initializeInput(inputPath string) (string, *os.File) {
	absPath, err := filepath.Abs(inputPath)
	checkError(err)

	file, err := os.Open(absPath)
	checkError(err)

	return absPath, file
}

func initializeOutput(inputFile string) *bufio.Writer {
	path, fileName := filepath.Split(inputFile)
	basename := strings.Split(fileName, ".")[0]
	outputFileName := fmt.Sprintf("%s%s.hack", path, basename)

	output, err := os.Create(outputFileName)
	checkError(err)

	return bufio.NewWriter(output)
}
