package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type Seat struct {
	Raw string
	Row int
	Col int
	ID  int
}

func ConvertSeat(raw string) Seat {
	// Row
	min := 0
	max := 127

	for i := 0; i < 7; i++ {
		c := raw[i]

		if c == 'F' {
			max -= (max - min + 1) / 2
		} else {
			min += (max - min + 1) / 2
		}
	}

	row := min

	// Col
	min = 0
	max = 7

	for i := 7; i < 10; i++ {
		c := raw[i]

		if c == 'L' {
			max -= (max - min + 1) / 2
		} else {
			min += (max - min + 1) / 2
		}
	}

	col := min

	return Seat{
		Raw: raw,
		Row: row,
		Col: col,
		ID:  row*8 + col,
	}
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	// Read seats
	var seats []Seat

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		seat := ConvertSeat(line)
		seats = append(seats, seat)
	}

	// Part one, look for highest seat
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].ID > seats[j].ID
	})

	fmt.Printf("part one: highest ID: %d\n", seats[0].ID)

	// Part two, look for my seat
	// Not in the front or back row
	ids := make(map[int]bool)
	for _, seat := range seats {
		ids[seat.ID] = true
	}

	for _, seat := range seats {
		if seat.Row == 0 || seat.Row == 127 {
			continue
		}

		// Looking for missing number before/after
		if !ids[seat.ID-1] {
			fmt.Printf("ID %d missing\n", seat.ID-1)
		}

		if !ids[seat.ID+1] {
			fmt.Printf("ID %d missing\n", seat.ID+1)
		}
	}
}
