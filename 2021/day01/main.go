package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read raw measurements from a file.
	data, err := os.ReadFile("simple.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert measurements
	var measurements []int
	for _, value := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(value) == "" {
			continue
		}

		measurement, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}

		measurements = append(measurements, measurement)
	}

	// Part one
	// Count number of increases
	var count int
	if len(measurements) > 1 {
		for i := 1; i < len(measurements); i++ {
			if measurements[i-1] < measurements[i] {
				count++
			}
		}
	}

	fmt.Printf("%d increases\n", count)

	// Part two
	// Sliding window

	// First, compute the sums
	var sums []int
	for i := 0; i < len(measurements)-2; i++ {
		var sum int
		for j := 0; j < 3; j++ {
			sum += measurements[i+j]
		}
		sums = append(sums, sum)
	}

	fmt.Printf("sums: %v\n", sums)

	// Then, count number of increases
	count = 0
	if len(sums) > 1 {
		for i := 1; i < len(sums); i++ {
			if sums[i-1] < sums[i] {
				count++
			}
		}
	}

	fmt.Printf("%d sum increases\n", count)
}
