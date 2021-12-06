package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	GROUND = '.'
	TREE   = '#'
)

func traverse(mmap []string, right int, down int) int {
	var row, col, trees int
	width := len(mmap[0])

	for row < len(mmap) {
		if mmap[row][col] == TREE {
			trees++
		}

		row += down
		col = (col + right) % width
	}

	return trees
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	// Sanitize the matrix
	var mmap []string
	var width int
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if width == 0 {
			width = len(line)
		} else {
			if len(line) != width {
				log.Fatalf("incorrect line width, expected %d, got %d for %q", width, len(line), line)
			}
		}

		mmap = append(mmap, line)
	}

	fmt.Printf("map dimension: width=%d height=%d\n", width, len(mmap))

	// Part one
	// right 3, down 1
	fmt.Printf("part 1 (right 3, down 1): %d trees encountered\n", traverse(mmap, 3, 1))

	// Part two, multiple slopes
	var solution = 0
	for _, slope := range []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		trees := traverse(mmap, slope.right, slope.down)
		if solution == 0 {
			solution = trees
		} else {
			solution *= trees
		}
	}

	fmt.Printf("part 2 (multiple slopes): %d\n", solution)
}
