package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read instructions from file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	type instruction struct {
		order string
		value int
	}

	// Convert into strings
	var instructions []instruction
	for _, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		elements := strings.Split(line, " ")
		if len(elements) != 2 {
			log.Fatal("incorrect number of elements")
		}

		value, err := strconv.Atoi(elements[1])
		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, instruction{
			order: elements[0],
			value: value,
		})
	}

	// Initial position
	var x int
	var depth int
	var aim int

	// Move
	for _, instr := range instructions {
		switch instr.order {
		case "forward":
			x += instr.value
			depth += aim * instr.value
		case "down":
			aim += instr.value
		case "up":
			aim -= instr.value
		}
	}

	// Final result
	fmt.Printf("x=%d depth=%d result=%d\n", x, depth, x*depth)
}
