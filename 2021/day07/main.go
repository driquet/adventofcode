package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert positions
	var positions []int
	for _, rawPosition := range strings.Split(strings.TrimSpace(string(data)), ",") {
		position, err := strconv.Atoi(rawPosition)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, position)
	}

	// Find median
	sort.Ints(positions)
	median := positions[len(positions)/2]

	// Compute solution (part 1)
	var fuel int
	for _, position := range positions {
		diff := median - position
		if diff < 0 {
			diff = -diff
		}

		fuel += diff
	}

	fmt.Printf("sum fuel part one: %d\n", fuel)

	// Compute solution (part 2)
	// Brute force
	minValue := positions[0]
	maxValue := positions[len(positions)-1]

	scores := make(map[int]int)

	for currentPosition := minValue; currentPosition <= maxValue; currentPosition++ {
		var fuel int

		for _, position := range positions {
			// Compute fuel consumption for this position
			var min, max int
			if position < currentPosition {
				min = position
				max = currentPosition
			} else {
				min = currentPosition
				max = position
			}

			var addedCost int
			var currentFuel int
			for i := min; i < max; i++ {
				currentFuel += 1 + addedCost
				addedCost++
			}

			fmt.Printf("position %d: fuel consumption %d\n", position, currentFuel)
			fuel += currentFuel
		}

		scores[currentPosition] = fuel
		fmt.Printf("possible location: %d total fuel consumption %d\n\n", currentPosition, fuel)
	}

	// Look for cheapest consumption
	cheapestPosition := -1

	for position, consumption := range scores {
		if cheapestPosition == -1 || scores[cheapestPosition] > consumption {
			cheapestPosition = position
		}
	}

	fmt.Printf("sum fuel part two: position %d total fuel consumption %d\n", cheapestPosition, scores[cheapestPosition])

}
