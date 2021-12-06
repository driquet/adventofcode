package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read binary numbers from file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		numbers = append(numbers, strings.TrimSpace(line))
	}

	// Part 1
	var gamma string
	var epsilon string
	size := len(numbers[0])

	for i := 0; i < size; i++ {
		// Find the most common value at the i-th position
		bits := make(map[byte]int)
		for _, number := range numbers {
			bits[number[i]]++
		}
		ones := bits['1']
		zeros := bits['0']

		if zeros > ones {
			// Most common is zero
			gamma += "0"
			epsilon += "1"
		} else {
			// Most common is zero
			gamma += "1"
			epsilon += "0"
		}
	}

	// Part 2
	var oxygen string
	var co2 string

	oxygenValues := numbers[:]
	co2Values := numbers[:]

	// Oxygen first
	i := 0
	for i < size {
		fmt.Printf("oxygen:\n")
		fmt.Printf("\ti: %d\n", i)
		fmt.Printf("\tvalues : %d (%v)\n", len(oxygenValues), oxygenValues)

		// Find the most common value at the i-th position
		bits := make(map[byte]int)
		for _, number := range oxygenValues {
			bits[number[i]]++
		}
		ones := bits['1']
		zeros := bits['0']

		var mostCommon byte
		if ones >= zeros {
			mostCommon = '1'
		} else {
			mostCommon = '0'
		}

		// Filter values with the most common value at the i-th position
		var values []string
		for _, number := range oxygenValues {
			if number[i] == mostCommon {
				values = append(values, number)
			}
		}
		oxygenValues = values

		// Checking if there is only one value left
		if len(oxygenValues) == 1 {
			oxygen = oxygenValues[0]
			break
		}

		i++
	}

	// CO2 at last
	i = 0
	for i < size {
		fmt.Printf("co2:\n")
		fmt.Printf("\ti: %d\n", i)
		fmt.Printf("\tvalues : %d (%v)\n", len(co2Values), co2Values)

		// Find the most common value at the i-th position
		bits := make(map[byte]int)
		for _, number := range co2Values {
			bits[number[i]]++
		}
		ones := bits['1']
		zeros := bits['0']

		var leastCommon byte
		if zeros <= ones {
			leastCommon = '0'
		} else {
			leastCommon = '1'
		}

		// Filter values with the least common value at the i-th position
		var values []string
		for _, number := range co2Values {
			if number[i] == leastCommon {
				values = append(values, number)
			}
		}
		co2Values = values

		// Checking if there is only one value left
		if len(co2Values) == 1 {
			co2 = co2Values[0]
			break
		}

		i++
	}

	// Converting to int
	gammaValue, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilonValue, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	oxygenValue, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	co2Value, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("gamma: %s (%d)\n", gamma, gammaValue)
	fmt.Printf("epsilon: %s (%d)\n", epsilon, epsilonValue)
	fmt.Printf("oxygen: %s (%d)\n", oxygen, oxygenValue)
	fmt.Printf("co2: %s (%d)\n", co2, co2Value)
	fmt.Printf("result part 1: %d\n", gammaValue*epsilonValue)
	fmt.Printf("result part 2: %d\n", oxygenValue*co2Value)
}
