package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	row, col int
}

var instructionRe = regexp.MustCompile(`fold along (x|y)=(\d+)`)

func displayPoints(points map[point]bool) {
	// Look for dimensions
	var width, height int

	for p := range points {
		if p.col > width {
			width = p.col
		}
		if p.row > height {
			height = p.row
		}
	}

	// Create grid
	var grid [][]bool
	for row := 0; row <= height; row++ {
		grid = append(grid, make([]bool, width+1))
	}

	// Put points in the grid
	for p := range points {
		grid[p.row][p.col] = true
	}

	// Display grid
	for _, row := range grid {
		for _, p := range row {
			if p {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := make(map[point]bool)
	var instructions []point
	var pointsDone bool

	// Read line by line
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		if line == "" {
			pointsDone = true
			continue
		}

		// Points
		if !pointsDone {
			elts := strings.Split(line, ",")
			if len(elts) != 2 {
				log.Fatal("unexpected line")
			}

			col, err := strconv.Atoi(elts[0])
			if err != nil {
				log.Fatal("cannot convert col number")
			}

			row, err := strconv.Atoi(elts[1])
			if err != nil {
				log.Fatal("cannot convert row number")
			}

			points[point{
				row: row,
				col: col,
			}] = true

			continue
		}

		// Instructions
		groups := instructionRe.FindStringSubmatch(line)
		if len(groups) != 3 {
			log.Fatal("incorrect instruction")
		}

		var p point
		number, _ := strconv.Atoi(groups[2])
		if groups[1] == "x" {
			p.col = number
		} else {
			p.row = number
		}

		instructions = append(instructions, p)
	}

	for _, instruction := range instructions {
		fmt.Printf("points before: %d\n", len(points))
		displayPoints(points)

		if instruction.row != 0 {
			fmt.Printf("fold horizontally y=%d\n", instruction.row)
		} else {
			fmt.Printf("fold vertically=%d\n", instruction.col)
		}

		next := make(map[point]bool)

		for p := range points {
			np := p

			if instruction.row > 0 && np.row > instruction.row {
				// Horizontal fold
				fmt.Printf("flip point: %+v\n", np)
				np.row = instruction.row - (np.row - instruction.row)
				fmt.Printf("result: %+v\n", np)
			} else if instruction.col > 0 && np.col > instruction.col {
				// Vertical fold
				fmt.Printf("flip point: %+v\n", np)
				np.col = instruction.col - (np.col - instruction.col)
			}

			next[np] = true
		}

		points = next
		displayPoints(points)
		fmt.Printf("points after: %d\n\n", len(points))
	}

}
