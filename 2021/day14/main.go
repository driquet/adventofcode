package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read line by line
	r := bufio.NewReader(f)

	// First line, template
	template, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("missing line")
	}
	template = strings.TrimSpace(template)

	// Empty line
	_, err = r.ReadString('\n')
	if err != nil {
		log.Fatal("missing line")
	}

	// Insertion rules
	insertions := make(map[string]byte)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		if len(line) != 7 {
			fmt.Printf("line: %s (%d)\n", line, len(line))
			log.Fatal("incorrect insertion")
		}
		insertions[line[0:2]] = line[6]
	}

	// Break the pairs
	firstPair := template[0:2]
	lastPair := template[len(template)-2 : len(template)]

	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	fmt.Printf("pairs: %+v\n", pairs)
	fmt.Printf("first pair: %s\n", firstPair)
	fmt.Printf("last pair: %s\n", lastPair)

	for step := 0; step < 40; step++ {
		fmt.Printf("step %d\n", step)
		newPairs := make(map[string]int)

		for pair, count := range pairs {
			c, found := insertions[pair]
			if !found {
				newPairs[pair] += count
				continue
			}

			// Two new pairs
			newPairs[fmt.Sprintf("%c%c", pair[0], c)] += count
			newPairs[fmt.Sprintf("%c%c", c, pair[1])] += count

			if pair == firstPair {
				firstPair = fmt.Sprintf("%c%c", pair[0], c)
			}
			if pair == lastPair {
				lastPair = fmt.Sprintf("%c%c", c, pair[1])
			}
		}

		pairs = newPairs
		fmt.Printf("pairs: %+v\n", pairs)
		fmt.Printf("first pair: %s\n", firstPair)
		fmt.Printf("last pair: %s\n", lastPair)
		fmt.Printf("\n")
	}

	fmt.Printf("pairs: %+v\n", pairs)
	fmt.Printf("first pair: %s\n", firstPair)
	fmt.Printf("last pair: %s\n", lastPair)

	// Frequencies
	frequencies := make(map[byte]int)
	for pair, count := range pairs {
		// If first pair or other, count only first letter
		// The second one will be counted by another pair
		frequencies[pair[0]] += count

		// Except for the last pair
		if pair == lastPair {
			frequencies[pair[1]] += 1
		}
	}

	var min, max byte
	for b, count := range frequencies {
		if min == 0 || max == 0 {
			min = b
			max = b
		}

		if frequencies[min] > count {
			min = b
		}

		if frequencies[max] < count {
			max = b
		}
	}

	leastCommonCount := frequencies[min]
	mostCommonCount := frequencies[max]
	fmt.Printf("least common: %c %d\n", min, leastCommonCount)
	fmt.Printf("most common : %c %d\n", max, mostCommonCount)
	fmt.Printf("part one: %d\n", mostCommonCount-leastCommonCount)
}
