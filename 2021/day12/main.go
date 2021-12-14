package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	start = "start"
	end   = "end"
)

func smallCave(name string) bool {
	return strings.ToLower(name) == name
}

// Only one small cave can be visited twice
func canVisitSmallCave(path []string, name string) bool {
	visits := make(map[string]int)

	// Count visits per small cave
	for _, location := range path {
		if !smallCave(location) || location == start {
			continue
		}

		visits[location]++
	}

	if len(visits) == 0 {
		// No small cave yet
		return true
	}

	for cave, count := range visits {
		if count >= 2 {
			if cave != name {
				// Another cave already has been visited twice.
				// So, to visit "name", it should not have been visited so far.
				return visits[name] == 0
			}

			// The cave already has been visited twice
			return false
		}
	}

	return true
}

func inList(l []string, str string) bool {
	for _, elt := range l {
		if elt == str {
			return true
		}
	}
	return false
}

func insertHead(l []string, h string) []string {
	return append([]string{h}, l...)
}

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	connections := make(map[string]map[string]bool)

	// Read line by line
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		idx := strings.IndexByte(line, '-')
		a := line[0:idx]
		b := line[idx+1 : len(line)-1]

		if _, found := connections[a]; !found {
			connections[a] = make(map[string]bool)
		}

		if _, found := connections[b]; !found {
			connections[b] = make(map[string]bool)
		}

		connections[a][b] = true
		connections[b][a] = true
	}

	fmt.Printf("connections: %+v\n", connections)

	paths := explore(connections, []string{start})
	sort.Strings(paths)
	fmt.Printf("possible paths (%d):\n", len(paths))
	// for _, path := range paths {
	// 	fmt.Printf("\t%s\n", path)
	// }
}

func explore(connections map[string]map[string]bool, path []string) []string {
	location := path[len(path)-1]

	// Check if we arrived at the end
	if location == end {
		fmt.Printf("possible path: %+v\n", path)
		return []string{strings.Join(path, ",")}
	}

	var res []string
	for next := range connections[location] {
		if next == start {
			continue
		}

		if next != end && smallCave(next) && !canVisitSmallCave(path, next) {
			fmt.Printf("path=%v cannot explore small cave %s\n", path, next)
			continue
		}

		fmt.Printf("path=%v explore %s\n", path, next)
		res = append(res, explore(connections, append(path, next))...)
	}

	return res
}
