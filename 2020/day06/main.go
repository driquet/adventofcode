package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Group struct {
	Count  int
	Values map[string]int
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	var groups []*Group
	group := &Group{
		Values: make(map[string]int),
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			groups = append(groups, group)
			group = &Group{
				Values: make(map[string]int),
			}
			continue
		}

		for _, r := range line {
			group.Values[string(r)]++
		}

		group.Count++
	}

	// Part one, sum of all different choices for every groups
	var sum int
	for _, group := range groups {
		sum += len(group.Values)
	}

	fmt.Printf("part one sum: %d\n", sum)

	// Part two, sum of common choices for every groups
	sum = 0
	for _, group := range groups {
		for _, count := range group.Values {
			if count == group.Count {
				sum++
			}
		}
	}

	fmt.Printf("part two sum: %d\n", sum)
}
