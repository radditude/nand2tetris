package main

import (
	"fmt"
	"strconv"
	"strings"
)

func handleAInstruction(instruction string) string {
	debug(fmt.Sprintf("the a instruction is: %s", instruction))

	address := getAddress(instruction)
	intAddress, err := parseIntegerAddress(address)

	// a parse error means this is a symbol and not a number
	if err != nil {
		intAddress = getSymbol(address)
	}

	return fmt.Sprintf("0%015b", intAddress)
}

func isAInstruction(instruction string) bool {
	return string(instruction[0]) == "@"
}

// get the address string
func getAddress(instruction string) string {
	return strings.Replace(instruction, "@", "", 1)
}

// parse an address string as an integer
func parseIntegerAddress(address string) (int64, error) {
	return strconv.ParseInt(address, 10, 16)
}
