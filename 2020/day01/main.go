package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read file
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("cannot read file")
	}

	var values []int
	for _, value := range strings.Split(string(data), "\n") {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}

		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("cannot convert %s", value)
		}

		values = append(values, i)
	}

	// Part one: sum of two elements equals 2020, solution is multiplication
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			sum := values[i] + values[j]
			if sum == 2020 {
				fmt.Printf("solution 1: %d + %d = 2020 so flag is %d\n",
					values[i],
					values[j],
					values[i]*values[j],
				)
			}
		}
	}

	// Part two: sum of three elements equals 2020, solution is multiplication
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				sum := values[i] + values[j] + values[k]
				if sum == 2020 {
					fmt.Printf("solution 1: %d + %d + %d = 2020 so flag is %d\n",
						values[i],
						values[j],
						values[k],
						values[i]*values[j]*values[k],
					)
				}
			}
		}
	}
}
