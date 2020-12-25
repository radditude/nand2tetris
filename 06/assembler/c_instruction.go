package main

import (
	"fmt"
	"regexp"
	"strings"
)

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

func handleCInstruction(instruction string) string {
	debug(fmt.Sprintf("the c instruction is: %s", instruction))

	dest, comp, jump := splitInstruction(instruction)

	return fmt.Sprintf("111%s%s%s", getComp(comp), getDest(dest), getJump(jump))
}

func splitInstruction(instruction string) (string, string, string) {
	dest, comp, jump := "", "", ""
	hasJump := strings.Contains(instruction, ";")
	hasDest := strings.Contains(instruction, "=")
	regex := regexp.MustCompile(`=|;`)

	if hasJump && hasDest {
		split := regex.Split(instruction, 3)
		dest = split[0]
		comp = split[1]
		jump = split[2]
	} else if hasJump {
		split := regex.Split(instruction, 2)
		comp = split[0]
		jump = split[1]
	} else if hasDest {
		split := regex.Split(instruction, 2)
		dest = split[0]
		comp = split[1]
	}

	return dest, comp, jump
}
