package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	var joltages []int

	// Read joltages
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		joltage, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("unable to convert joltage")
		}
		joltages = append(joltages, joltage)
	}

	sort.Ints(joltages)

	// Part one, find the chain and count the differences
	differences := make(map[int]int)

	currentJoltage := 0

	for _, joltage := range joltages {
		difference := joltage - currentJoltage
		if difference > 3 {
			log.Fatal("something wrong with joltage")
		}

		differences[difference]++
		currentJoltage += difference
	}

	currentJoltage += 3
	differences[3]++

	fmt.Printf("current joltage: %d\n", currentJoltage)
	fmt.Printf("differences: %v\n", differences)

	fmt.Printf("part one solution: %d\n", differences[1]*differences[3])

	// Part two solution, find combinations

	fmt.Printf("\n\n##############\n\n")
	combinations := findCombination(0, joltages)

	fmt.Printf("%d combinations\n", combinations)
}

func findCombination(currentJoltage int, joltages []int) int {
	var res int

	if len(joltages) == 1 {
		return 1
	}

	for i := 0; i < 3 && i < len(joltages) && (joltages[i]-currentJoltage) <= 3; i++ {
		subJoltages := joltages[i+1:]
		if len(subJoltages) > 0 {
			res += findCombination(joltages[i], subJoltages)
		} else {
			res += 1
		}
	}

	return res
}
