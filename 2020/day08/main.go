package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Argument  int
}

var instructionRe = regexp.MustCompile(`(nop|acc|jmp) ((\+|-)\d+)`)

func runInstruction(instructions []*Instruction) (int, bool) {
	var accumulator int
	var idx int
	visited := make(map[int]bool)

	for {
		// Detecting loop
		if _, found := visited[idx]; found {
			return accumulator, false
		}
		visited[idx] = true

		instr := instructions[idx]

		switch instr.Operation {
		case "nop":
			idx++

		case "acc":
			accumulator += instr.Argument
			idx++

		case "jmp":
			idx += instr.Argument
		}

		// Dealing with termination
		if idx == len(instructions) {
			return accumulator, true
		}
	}

	return 0, false
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	// Convert instructions
	var instructions []*Instruction

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		match := instructionRe.FindStringSubmatch(line)
		if len(match) != 4 {
			fmt.Printf("cannot extract instruction: %q\n", line)
			continue
		}

		argument, _ := strconv.Atoi(match[2])

		instructions = append(instructions, &Instruction{
			Operation: match[1],
			Argument:  argument,
		})
	}

	// Part one, run the code until we reach a loop
	var accumulator int
	var idx int
	visited := make(map[int]bool)

	for {
		if _, found := visited[idx]; found {
			break
		}
		visited[idx] = true

		instr := instructions[idx]

		switch instr.Operation {
		case "nop":
			idx++

		case "acc":
			accumulator += instr.Argument
			idx++

		case "jmp":
			idx += instr.Argument
		}
	}

	fmt.Printf("part one: accumulator value is %d when loop detected\n", accumulator)

	// Part two, fix nop/jmp instructions
	for idx, instr := range instructions {
		if instr.Operation == "acc" {
			continue
		}

		// Store old operation
		oldOperation := instr.Operation

		// Replace operation
		if oldOperation == "nop" {
			instr.Operation = "jmp"
		} else {
			instr.Operation = "nop"
		}

		fmt.Printf("trying to fix instruction %d (%s)\n", idx, oldOperation)

		acc, valid := runInstruction(instructions)
		if valid {
			fmt.Printf("part two: fix instruction %d, accumulator is %d\n", idx, acc)
		}

		// Reset operation
		instr.Operation = oldOperation
	}
}
