package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	X, Y int
}

type line struct {
	from point
	to   point
}

func convertInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func computePoints(line line) []point {
	var points []point

	if line.from.X == line.to.X {
		// Vertical line
		var min, max int
		if line.from.Y < line.to.Y {
			min = line.from.Y
			max = line.to.Y
		} else {
			min = line.to.Y
			max = line.from.Y
		}

		for y := min; y <= max; y++ {
			points = append(points, point{
				X: line.from.X,
				Y: y,
			})
		}
	} else if line.from.Y == line.to.Y {
		// Horizontal line
		var min, max int
		if line.from.X < line.to.X {
			min = line.from.X
			max = line.to.X
		} else {
			min = line.to.X
			max = line.from.X
		}

		for x := min; x <= max; x++ {
			points = append(points, point{
				X: x,
				Y: line.from.Y,
			})
		}
	} else {
		fmt.Printf("diagonal\n")
		// Diagonal line at exactly 45°
		// Determine the point closest to the left
		var closest, farthest point
		if line.from.X < line.to.X {
			closest = line.from
			farthest = line.to
		} else {
			closest = line.to
			farthest = line.from
		}

		// From the closest point, it either goes up or down
		var step int
		if closest.Y < farthest.Y {
			step = 1
		} else {
			step = -1
		}

		y := closest.Y

		for x := closest.X; x <= farthest.X; x++ {
			points = append(points, point{
				X: x,
				Y: y,
			})

			y += step
		}

	}

	return points
}

func displayDiagram(diagram [][]int) {
	for _, row := range diagram {
		for i := 0; i < len(row); i++ {
			if row[i] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", row[i])
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

	// Create reader
	r := bufio.NewReader(f)

	// Convert lines
	var lines []line
	lineRe := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	for {
		raw, err := r.ReadString('\n')
		if err != nil {
			break
		}

		groups := lineRe.FindStringSubmatch(raw)
		if len(groups) != 5 {
			log.Fatal("cannot convert line")
		}

		lines = append(lines, line{
			from: point{
				X: convertInt(groups[1]),
				Y: convertInt(groups[2]),
			},
			to: point{
				X: convertInt(groups[3]),
				Y: convertInt(groups[4]),
			},
		})
	}

	// Find max values
	max := point{}

	for _, line := range lines {
		if line.from.X > max.X {
			max.X = line.from.X
		}
		if line.to.X > max.X {
			max.X = line.to.X
		}
		if line.from.Y > max.Y {
			max.Y = line.from.Y
		}
		if line.to.Y > max.Y {
			max.Y = line.to.Y
		}
	}

	// Build the diagram
	var diagram [][]int
	for row := 0; row <= max.Y; row++ {
		diagram = append(diagram, make([]int, max.X+1))
	}

	// Add lines
	for _, line := range lines {
		fmt.Printf("line: %+v\n", line)
		for _, point := range computePoints(line) {
			fmt.Printf("point: %+v\n", point)
			diagram[point.Y][point.X]++
		}
	}

	displayDiagram(diagram)

	// Compute solution
	var dangers int

	for _, row := range diagram {
		for x := 0; x < len(row); x++ {
			if row[x] >= 2 {
				dangers++
			}
		}
	}

	fmt.Printf("dangers: %d\n", dangers)
}
