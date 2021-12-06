package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func findFirstInvalid(numbers []int, size int) (int, int) {
	for idx := size; idx < len(numbers); idx++ {
		fmt.Printf("idx %d: value %d\n", idx, numbers[idx])

		value := numbers[idx]

		var found bool
		for i := idx - size; i < idx; i++ {
			for j := idx - size + 1; j < idx; j++ {
				if numbers[i]+numbers[j] == value {
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		if !found {
			fmt.Printf("invalid value %d at index %d\n", value, idx)
			return idx, value
		}
	}

	return 0, 0
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	var numbers []int

	// Reading numbers
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("unable to convert number")
		}

		numbers = append(numbers, number)
	}

	// Part one, find the weakness
	idx, value := findFirstInvalid(numbers, 25)

	// Part two, find contiguous values that equal this weakness when added up
	var found bool
	for i := 0; i < idx; i++ {
		values := []int{numbers[i]}
		sum := numbers[i]

		for j := i + 1; j < idx; j++ {
			sum += numbers[j]
			values = append(values, numbers[j])

			if sum == value {
				fmt.Printf("found value (%d-%d)\n", i, j)
				found = true

				sort.Ints(values)
				fmt.Printf("part two: %d+%d=%d\n", values[0], values[len(values)-1], values[0]+values[len(values)-1])

				break
			}

			if sum > value {
				break
			}
		}

		if found {
			break
		}
	}
}
