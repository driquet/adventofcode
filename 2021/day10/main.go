package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type charHeap []rune

func (h *charHeap) push(r rune) {
	*h = append(*h, r)
}

func (h *charHeap) pop() (rune, bool) {
	var r rune

	old := *h

	if len(old) == 0 {
		return r, false
	}

	r = old[len(old)-1]
	*h = old[0 : len(old)-1]

	return r, true
}

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string

	// Read line by line
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, strings.TrimSpace(line))
	}

	// Part one: find corrupted lines
	var score int

	corruptedScorePerChar := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	incompleteScorePerChar := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	var incompleteScores []int

	for _, line := range lines {
		var corrupted bool
		var heap charHeap

		for _, r := range line {
			switch r {
			case '(':
				heap.push(')')
			case '[':
				heap.push(']')
			case '<':
				heap.push('>')
			case '{':
				heap.push('}')

			case ')', ']', '>', '}':
				expected, found := heap.pop()
				if !found {
					log.Fatal("empty heap")

					break
				}
				if r != expected {
					fmt.Printf("unexpected char %c, expected %c\n", r, expected)
					corrupted = true
					score += corruptedScorePerChar[r]
					break
				}

			default:
				log.Fatal("unexpected character")
			}
		}

		if len(heap) > 0 {
			// Incomplete line
			// Compute incomplete score
			var incompleteScore int
			for {
				expected, found := heap.pop()
				if !found {
					break
				}
				incompleteScore *= 5
				incompleteScore += incompleteScorePerChar[expected]

			}

			fmt.Printf("incomplete %s: score %d\n", line, incompleteScore)
			incompleteScores = append(incompleteScores, incompleteScore)

		}

		if corrupted {
			fmt.Printf("corrupted: %s\n", line)
		}
	}

	fmt.Printf("part one: %d\n", score)

	sort.Ints(incompleteScores)
	fmt.Printf("part two: %d\n", incompleteScores[len(incompleteScores)/2])
}
