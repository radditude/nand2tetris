package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var symbolTable = initializeSymbolTable()

var currentAddress int64 = 16

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

func scanForSymbols(input *os.File) {
	symbolScanner := bufio.NewScanner(input)
	var currentInstruction int64 = 0

	for symbolScanner.Scan() {
		line := symbolScanner.Text()

		// remove comments and whitespace
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
}

func getSymbol(address string) int64 {
	var intAddress int64

	if val, ok := symbolTable[address]; ok {
		intAddress = val
	} else {
		symbolTable[address] = currentAddress
		intAddress = currentAddress
		currentAddress++
	}

	return intAddress
}

func isLabel(instruction string) bool {
	return strings.Contains(instruction, "(") && strings.Contains(instruction, ")")
}
