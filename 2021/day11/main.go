package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	energyFlash = 9
)

type point struct {
	row, col int
}

func adjacent(p point, width, height int) []point {
	var res []point

	// Top
	if p.row > 0 {
		res = append(res, point{
			row: p.row - 1,
			col: p.col,
		})

		// Top left
		if p.col > 0 {
			res = append(res, point{
				row: p.row - 1,
				col: p.col - 1,
			})
		}

		// Top right
		if p.col < height-1 {
			res = append(res, point{
				row: p.row - 1,
				col: p.col + 1,
			})
		}
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

	// Bottom
	if p.row < height-1 {
		res = append(res, point{
			row: p.row + 1,
			col: p.col,
		})

		// Bottom left
		if p.col > 0 {
			res = append(res, point{
				row: p.row + 1,
				col: p.col - 1,
			})
		}

		// Bottom right
		if p.col < width-1 {
			res = append(res, point{
				row: p.row + 1,
				col: p.col + 1,
			})
		}
	}

	return res
}

func displayGrid(grid [][]int, width, height int) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if grid[row][col] > energyFlash {
				if grid[row][col] == energyFlash+1 {
					fmt.Printf("F")
				} else {
					fmt.Printf("X")
				}
			} else {
				fmt.Printf("%d", grid[row][col])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read the grid line by line
	var grid [][]int
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		var row []int
		for _, r := range strings.TrimSpace(line) {
			row = append(row, int(r)-'0')
		}

		grid = append(grid, row)
	}

	height := len(grid)
	width := len(grid[0])

	// Part one: how many flashes in X steps
	X := 500
	var flashes int
	var allFlashes []int

	for step := 0; step < X; step++ {
		fmt.Printf("just before step %d\n", step)
		displayGrid(grid, width, height)

		// First, the energy level of each octopus increases by 1
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				grid[row][col]++
			}
		}

		// Look for octopus with an energy level greater than 9
		var flashing []point
		flashed := make(map[point]bool)

		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				if grid[row][col] > energyFlash {
					p := point{
						row: row,
						col: col,
					}
					flashed[p] = true
					flashing = append(flashing, p)
				}
			}
		}

		fmt.Printf("%d octopus are flashing at beginning of step %d\n", len(flashing), step)
		displayGrid(grid, width, height)

		// Each flashing octopus increases the energy level of all adjacent octopuses.
		// This is recursive: if, in this step, another octopuses gets to be flashing, it will increases its adjacent
		// octopuses' energy level.
		for len(flashing) > 0 {

			// Increase energy level of adjacent octopuses
			for _, p := range flashing {
				flashed[p] = true

				for _, adj := range adjacent(p, width, height) {
					grid[adj.row][adj.col]++
				}
			}

			// Looking for newly flashing octopuses
			var next []point
			for row := 0; row < height; row++ {
				for col := 0; col < width; col++ {
					p := point{
						row: row,
						col: col,
					}
					if grid[row][col] > energyFlash && !flashed[p] {
						fmt.Printf("newly flashing octopus: %+v\n", p)
						next = append(next, p)
					}
				}
			}

			flashing = next

			fmt.Printf("next iteration\n")
			displayGrid(grid, width, height)
		}

		// Count the number of flashes
		var count int

		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				if grid[row][col] > energyFlash {
					count++
					grid[row][col] = 0
				}
			}
		}

		fmt.Printf("step %d: %d flashes\n", step, count)
		flashes += count

		fmt.Printf("after step %d\n", step)
		displayGrid(grid, width, height)

		if count == width*height {
			fmt.Printf("all flashes!\n")
			allFlashes = append(allFlashes, step)
		}

	}

	fmt.Printf("total flashes: %d\n", flashes)
	fmt.Printf("all flashes: %+v\n", allFlashes)
}
