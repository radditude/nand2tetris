package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var debugEnabled bool = false

var jumpTable = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

var destTable = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

var compTable = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

var symbolTable = initializeSymbolTable()

var currentAddress int64 = 16

func main() {
	argError := checkArgs()

	if argError != "" {
		fmt.Println(argError)
		return
	}

	// get the name of the asm file
	inputPath, err := filepath.Abs(os.Args[1])
	checkError(err)

	// make sure the input file exists
	inputFile, err := os.Open(inputPath)
	defer inputFile.Close()
	checkError(err)

	// create and open the output file
	outputFile := prepareOutputFile(inputPath)
	output := bufio.NewWriter(outputFile)
	defer output.Flush()

	// first pass to identify labels
	symbolScanner := bufio.NewScanner(inputFile)
	var currentInstruction int64 = 0

	for symbolScanner.Scan() {
		line := symbolScanner.Text()

		// remove comments
		line = ignoreComments(line)

		if line != "" {
			if isLabel(line) {
				symbol := strings.ReplaceAll(line, "(", "")
				symbol = strings.ReplaceAll(symbol, ")", "")
				symbolTable[symbol] = currentInstruction
			} else {
				currentInstruction++
			}
		}
	}

	if err := symbolScanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
		return
	}

	// second pass to actually assemble
	inputFile.Seek(0, 0)
	parseScanner := bufio.NewScanner(inputFile)

	for parseScanner.Scan() {
		line := parseScanner.Text()
		var outputLine string

		// remove comments
		line = ignoreComments(line)
		// ignore labels
		line = ignoreLabels(line)

		if line != "" {
			if isAInstruction(line) {
				outputLine = handleAInstruction(line)
			} else {
				outputLine = handleCInstruction(line)
			}

			fmt.Fprintln(output, outputLine)
		}
	}

	fmt.Printf("done! 😎\n")
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

func prepareOutputFile(inputFile string) *os.File {
	path, fileName := filepath.Split(inputFile)
	basename := strings.Split(fileName, ".")[0]
	outputFileName := fmt.Sprintf("%s%s.hack", path, basename)
	output, err := os.Create(outputFileName)
	checkError(err)

	return output
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

func isAInstruction(instruction string) bool {
	return string(instruction[0]) == "@"
}

func isLabel(instruction string) bool {
	return strings.Contains(instruction, "(") && strings.Contains(instruction, ")")
}

func getAddress(instruction string) string {
	return strings.Replace(instruction, "@", "", 1)
}

func parseIntegerAddress(address string) (int64, error) {
	return strconv.ParseInt(address, 10, 16)
}

func handleAInstruction(instruction string) string {
	debug(fmt.Sprintf("the a instruction is: %s", instruction))

	address := getAddress(instruction)
	debug(fmt.Sprintf("the address is: %s", address))
	intAddress, err := parseIntegerAddress(address)

	// a parse error means this is a symbol and not a number
	if err != nil {
		if val, ok := symbolTable[address]; ok {
			debug(fmt.Sprintf("the retrieved symbol value is: %d\n", val))
			intAddress = val
		} else {
			debug(fmt.Sprintf("setting the symbol value to: %d\n", currentAddress))
			symbolTable[address] = currentAddress
			intAddress = currentAddress
			currentAddress++
		}
	}

	return fmt.Sprintf("0%015b", intAddress)
}

func handleCInstruction(instruction string) string {
	debug(fmt.Sprintf("the c instruction is: %s", instruction))
	var dest string
	var comp string
	var jump string
	firstSlice := strings.Split(instruction, "=")
	var secondSlice []string

	if len(firstSlice) == 1 {
		dest = ""
		secondSlice = strings.Split(firstSlice[0], ";")
	} else {
		dest = firstSlice[0]
		secondSlice = strings.Split(firstSlice[1], ";")
	}

	comp = secondSlice[0]

	if len(secondSlice) == 1 {
		jump = ""
	} else {
		jump = secondSlice[1]
	}

	return fmt.Sprintf("111%s%s%s", getComp(comp), getDest(dest), getJump(jump))
}

func getDest(dest string) string {
	debug(fmt.Sprintf("the destination is: %s", dest))

	if val, ok := destTable[dest]; ok {
		return val
	}

	panic("destination not found")
}

func getJump(jump string) string {
	debug(fmt.Sprintf("the jump is: %s", jump))
	if val, ok := jumpTable[jump]; ok {
		return val
	}

	panic("jump not found")
}

func getComp(comp string) string {
	debug(fmt.Sprintf("the comp is: %s", comp))

	if val, ok := compTable[comp]; ok {
		return val
	}

	panic("comp not found")
}

func initializeSymbolTable() map[string]int64 {
	table := map[string]int64{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	for i := 0; i < 16; i++ {
		key := fmt.Sprintf("R%d", i)
		table[key] = int64(i)
	}

	return table
}
