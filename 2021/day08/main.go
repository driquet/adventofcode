package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entry struct {
	signals []string
	outputs []string
}

func sortStringByCharacter(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func sortStrings(values []string) []string {
	var res []string

	for _, value := range values {
		res = append(res, sortStringByCharacter(value))
	}

	return res
}

func contains(str string, sub string) bool {
	for _, c := range sub {
		if !strings.ContainsRune(str, c) {
			return false
		}
	}
	return true
}

func minus(str string, sub string) string {
	var res []rune
	for _, c := range str {
		if !strings.ContainsRune(sub, c) {
			res = append(res, c)
		}
	}
	return string(res)
}

func main() {
	// Read input from file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var entries []entry

	// Read input line by line
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatal("incorrect line")
		}

		entries = append(entries, entry{
			signals: sortStrings(strings.Fields(parts[0])),
			outputs: sortStrings(strings.Fields(parts[1])),
		})
	}

	digitPerSize := map[int][]int{
		2: {1},
		3: {7},
		4: {4},
		5: {2, 3, 5},
		6: {0, 6, 9},
		7: {8},
	}

	// Part one: find digits with unique features
	var sum int
	for _, entry := range entries {
		var count int

		// Count known digits in output
		for _, output := range entry.outputs {
			// Look for output whose len is associated to only one digit
			digits := digitPerSize[len(output)]
			if len(digits) == 1 {
				count++
			}
		}

		fmt.Printf("%d simple digits\n", count)
		sum += count
	}

	fmt.Printf("part one: %d\n", sum)

	// Part two: find the associations
	sum = 0
	for _, entry := range entries {
		fmt.Printf("entry: %+v\n", entry)

		associationsPerSignal := make(map[string]int)
		associationsPerDigit := make(map[int]string)

		// Find simple associations: 1, 4, 7 and 8
		for _, signal := range entry.signals {
			digits := digitPerSize[len(signal)]
			if len(digits) == 1 {
				associationsPerSignal[signal] = digits[0]
				associationsPerDigit[digits[0]] = signal
			}
		}

		// Signals with 5 letters: 2, 3, 5
		var fives []string
		for _, signal := range entry.signals {
			if len(signal) == 5 {
				fives = append(fives, signal)
			}
		}

		fmt.Printf("fives: %v\n", fives)

		// Determine first which is 3: 3 is composed of 1 while 2 and 5 are not
		for _, five := range fives {
			if contains(five, associationsPerDigit[1]) {
				associationsPerSignal[five] = 3
				associationsPerDigit[3] = five
			}
		}

		// Then, between 2 and 5, we can determine 5 thanks to 4
		// As we know 1 and 4, we now the signal associated to b and d (digit 4 minus digit 1)
		// b and d are used in 5 and not in 2
		bd := minus(associationsPerDigit[4], associationsPerDigit[1])
		for _, five := range fives {
			if contains(five, bd) {
				associationsPerSignal[five] = 5
				associationsPerDigit[5] = five
			}
		}

		// Last one is 2
		for _, five := range fives {
			if _, found := associationsPerSignal[five]; !found {
				associationsPerSignal[five] = 2
				associationsPerDigit[2] = five
			}
		}

		// Signals with 6 letters: 0, 6, 9
		var sixes []string
		for _, signal := range entry.signals {
			if len(signal) == 6 {
				sixes = append(sixes, signal)
			}
		}

		// First, determine 6, because 6 is the only one that does not contain one
		for _, six := range sixes {
			if !contains(six, associationsPerDigit[1]) {
				associationsPerSignal[six] = 6
				associationsPerDigit[6] = six
			}
		}

		// Then, determine 9 which is the only one containing 4
		for _, six := range sixes {
			if contains(six, associationsPerDigit[4]) {
				associationsPerSignal[six] = 9
				associationsPerDigit[9] = six
			}
		}

		// Last is 0
		for _, six := range sixes {
			if _, found := associationsPerSignal[six]; !found {
				associationsPerSignal[six] = 0
				associationsPerDigit[0] = six
			}
		}

		fmt.Printf("associations per signal: %+v\n", associationsPerSignal)
		fmt.Printf("associations per digit: %+v\n", associationsPerDigit)

		// Determining the output
		var value string
		for _, output := range entry.outputs {
			value += fmt.Sprintf("%d", associationsPerSignal[output])
		}
		fmt.Printf("value: %s\n", value)

		val, _ := strconv.Atoi(value)
		sum += val
	}

	fmt.Printf("part two: %d\n", sum)
}
