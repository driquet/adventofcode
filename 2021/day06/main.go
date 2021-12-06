package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	reset   = 6
	newborn = 8
)

func main() {
	// Read input
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Store the number of lanternfish per age
	ages := make([]int, newborn+1)

	// Convert input into lanternfish ages
	for _, val := range strings.Split(strings.TrimSpace(string(data)), ",") {
		age, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		ages[age]++
	}

	// Part one
	count := 256

	for day := 0; day < count; day++ {
		// Print stuff
		for i := 0; i <= newborn; i++ {
			fmt.Printf("%d (%d) ", i, ages[i])
		}
		fmt.Printf("\n")

		// Store the number of lanternfish that will produce a newborn
		newborns := ages[0]

		// Shift values
		newAges := make([]int, newborn+1)
		for i := 0; i < newborn; i++ {
			newAges[i] = ages[i+1]
		}
		ages = newAges

		// Create newborns
		ages[newborn] = newborns
		// Reset the ones that gave birth
		ages[reset] += newborns

		// Count the number of fishes
		var total int
		for i := 0; i < len(ages); i++ {
			total += ages[i]
		}
		fmt.Printf("day %3d: %d\n", day+1, total)
	}

}
