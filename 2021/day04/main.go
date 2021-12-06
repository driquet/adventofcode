package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const size = 5

type cell struct {
	value  int
	marked bool
}

type board struct {
	cells [size][size]cell
}

func (b *board) mark(value int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if b.cells[y][x].value == value {
				b.cells[y][x].marked = true
				return
			}
		}
	}
}

func (b *board) sumUnmarked() int {
	var sum int
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if !b.cells[y][x].marked {
				sum += b.cells[y][x].value
			}
		}
	}

	return sum
}

func (b *board) display() {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Printf("%2d", b.cells[y][x].value)
			if b.cells[y][x].marked {
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

func (b *board) bingo() bool {
	// Check rows
	for y := 0; y < size; y++ {
		allMarked := true
		for x := 0; x < size; x++ {
			if !b.cells[y][x].marked {
				allMarked = false
				break
			}
		}

		if allMarked {
			return true
		}
	}

	// Check columns
	for x := 0; x < size; x++ {
		allMarked := true
		for y := 0; y < size; y++ {
			if !b.cells[y][x].marked {
				allMarked = false
				break
			}
		}

		if allMarked {
			return true
		}
	}

	return false
}

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read file line by line
	r := bufio.NewReader(f)

	// First line is the values
	rawValues, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var values []int
	for _, rawValue := range strings.Split(strings.TrimSpace(rawValues), ",") {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}

	fmt.Printf("values: %+v\n", values)

	var boards []*board
	// Then the boards
	for {
		// Read empty line
		_, err := r.ReadString('\n')
		if err != nil {
			// End of input
			break
		}

		// Create a new board
		b := &board{
			cells: [size][size]cell{},
		}

		for row := 0; row < size; row++ {
			// Read the line
			line, err := r.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			// Parse the values
			for col, rawValue := range strings.Fields(line) {
				value, err := strconv.Atoi(rawValue)
				if err != nil {
					log.Fatal(err)
				}

				// Set value
				b.cells[row][col].value = value
			}
		}

		// Append the board
		boards = append(boards, b)
	}

	// Just display the board for now
	for i, b := range boards {
		fmt.Printf("board %d\n", i)
		b.display()
		fmt.Printf("\n")
	}

	bingoed := make(map[int]bool)

	// Play bingo
	for round, value := range values {
		fmt.Printf("round %d: %d\n", round, value)

		// Mark values
		for _, b := range boards {
			if !b.bingo() {
				b.mark(value)
			}
		}

		// Check if we have a winner
		for idx, b := range boards {
			if bingoed[idx] {
				continue
			}

			if b.bingo() {
				sum := b.sumUnmarked()
				fmt.Printf("%d bingo\n", idx)
				fmt.Printf("result: sum=%d value=%d result=%d\n", sum, value, sum*value)
				bingoed[idx] = true
			}
		}
	}
}
