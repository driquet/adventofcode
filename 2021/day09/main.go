package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode"
)

type point struct {
	row, col int
}

func neighbours(p point, width, height int) []point {
	var res []point

	// Up
	if p.row > 0 {
		res = append(res, point{
			row: p.row - 1,
			col: p.col,
		})
	}

	// Down
	if p.row < height-1 {
		res = append(res, point{
			row: p.row + 1,
			col: p.col,
		})
	}

	// Left
	if p.col > 0 {
		res = append(res, point{
			row: p.row,
			col: p.col - 1,
		})
	}

	// Right
	if p.col < width-1 {
		res = append(res, point{
			row: p.row,
			col: p.col + 1,
		})
	}

	return res
}

func main() {
	// Read data from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read heightmap
	var hmap [][]int

	// Read line by line
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		var row []int
		for _, char := range line {
			if unicode.IsDigit(char) {
				row = append(row, int(char)-'0')
			}
		}

		hmap = append(hmap, row)
	}

	// Display map
	for _, row := range hmap {
		for _, value := range row {
			fmt.Printf("%d", value)
		}
		fmt.Printf("\n")
	}

	// Part one
	// Find low points
	width := len(hmap[0])
	height := len(hmap)
	var lowpoints []point
	var sum int

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			value := hmap[row][col]

			// Checking the four directions individually

			// Up
			if row > 0 && value >= hmap[row-1][col] {
				// Up value is lower or equal
				continue
			}

			// Down
			if row < height-1 && value >= hmap[row+1][col] {
				// Down value is lower or equal
				continue
			}

			// Left
			if col > 0 && value >= hmap[row][col-1] {
				// Left value is lower or equal
				continue
			}

			// Right
			if col < width-1 && value >= hmap[row][col+1] {
				// Right value is lower or equal
				continue
			}

			// This value is a lowpoint
			fmt.Printf("lowpoint (%dx%d): %d\n", col, row, value)
			sum += value + 1

			lowpoints = append(lowpoints, point{
				row: row,
				col: col,
			})
		}
	}

	fmt.Printf("part one: %d\n", sum)

	// Part two
	// Find basins
	var basinSizes []int

	for _, lowpoint := range lowpoints {
		// for each lowpoint, explore the new neighbors
		visited := make(map[point]bool)
		var basin []point

		toVisit := []point{lowpoint}
		for len(toVisit) > 0 {
			var nextToVisit []point

			for _, current := range toVisit {
				if visited[current] {
					continue
				}

				// Add current point to basin
				basin = append(basin, current)
				visited[current] = true

				// Add its neighbours to next visit
				for _, neighbour := range neighbours(current, width, height) {
					if hmap[neighbour.row][neighbour.col] != 9 {
						nextToVisit = append(nextToVisit, neighbour)
					}
				}
			}

			// Update visit list
			toVisit = nextToVisit
		}

		fmt.Printf("lowpoint: %+v\n", lowpoint)
		fmt.Printf("basin: %+v\n", basin)
		fmt.Printf("\n")

		basinSizes = append(basinSizes, len(basin))
	}

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})
	fmt.Printf("basin sizes: %+v\n", basinSizes)
	fmt.Printf("result: %d\n", basinSizes[0]*basinSizes[1]*basinSizes[2])
}
