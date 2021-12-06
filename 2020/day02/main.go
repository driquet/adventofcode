package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var lineRe = regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("cannot read file")
	}

	var validCount int

	// Part one:
	// A-B C: D
	// Range A to B
	// Char C must be between A to B times in D
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := lineRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalf("cannot match line: %q", line)
		}

		if len(matches) < 5 {
			log.Fatalf("not element number of elements")
		}

		from, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("incorrect from value: %q", matches[1])
		}

		to, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("incorrect to value: %q", matches[2])
		}

		count := strings.Count(matches[4], matches[3])
		if count < from || count > to {
			log.Printf("invalid password: %q", line)
		} else {
			validCount++
		}
	}

	fmt.Printf("part one valid count: %d\n", validCount)

	// Part Two:
	// A-B C: D
	// A and B are indexes (starting from 1)
	// Char C must be at one of the location A or B in D
	validCount = 0

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := lineRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalf("cannot match line: %q", line)
		}

		if len(matches) < 5 {
			log.Fatalf("not element number of elements")
		}

		i, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("incorrect i value: %q", matches[1])
		}

		j, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("incorrect j value: %q", matches[2])
		}

		c := matches[3][0]
		word := matches[4]

		i--
		j--

		var count int
		if i < len(word) && word[i] == c {
			count++
		}
		if j < len(word) && word[j] == c {
			count++
		}

		if count != 1 {
			log.Printf("invalid password: %q", line)
		} else {
			validCount++
		}
	}

	fmt.Printf("part two valid count: %d\n", validCount)
}
