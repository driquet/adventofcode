package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	lineFormatRe  = regexp.MustCompile(`(\w+ \w+) bags contain (.*)`)
	colorFormatRe = regexp.MustCompile(`(\d+) (\w+ \w+) bag`)
)

func computeBagCanContain(color string, canContain map[string]map[string]int) int {
	bagCanContain, found := canContain[color]
	if !found {
		return 1
	}

	sum := 1 // The initial bag
	for containedColor, count := range bagCanContain {
		sum += count * computeBagCanContain(containedColor, canContain)
	}

	return sum
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	// Store bag capacity to be hold by another bag
	bagCanBeHeldBy := make(map[string]map[string]bool)
	bagCanContain := make(map[string]map[string]int)

	// Part one, find how many combination of bags can hold a shiny bag
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Skip bags that cannot contain other bags
		if strings.Contains(line, "contain no other bags") {
			continue
		}

		// Line is formatted as follow:
		// [color] contain [[color] bags?]*
		// [color]: adj color
		match := lineFormatRe.FindStringSubmatch(line)
		if len(match) != 3 {
			fmt.Printf("line cannot be processed: %q\n", line)
			continue
		}

		color := match[1]
		containedColors := match[2]

		if _, found := bagCanContain[color]; !found {
			bagCanContain[color] = make(map[string]int)
		}

		// Find contained colors
		matches := colorFormatRe.FindAllStringSubmatch(containedColors, -1)
		for _, match := range matches {
			containedColor := match[2]

			if _, found := bagCanBeHeldBy[containedColor]; !found {
				bagCanBeHeldBy[containedColor] = make(map[string]bool)
			}

			number, _ := strconv.Atoi(match[1])

			bagCanBeHeldBy[containedColor][color] = true
			bagCanContain[color][containedColor] = number
		}
	}

	// Look for the combination
	colors := []string{"shiny gold"}
	outermostBags := make(map[string]bool)

	for len(colors) > 0 {
		var newColors []string

		for _, color := range colors {
			for canBeContainedBy := range bagCanBeHeldBy[color] {
				outermostBags[canBeContainedBy] = true
				newColors = append(newColors, canBeContainedBy)
			}
		}

		colors = newColors
	}

	fmt.Printf("part one: a shiny gold bag can be contained by %d bags\n", len(outermostBags))

	// Part two, look how many a bag can hold other bags
	fmt.Printf("part two: a shiny gold bag can contain %d bags\n", computeBagCanContain("shiny gold", bagCanContain)-1)

}
