package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestConvertSeat(t *testing.T) {
	for _, test := range []struct {
		input string
		seat  Seat
	}{
		{
			input: "FBFBBFFRLR",
			seat: Seat{
				Raw: "FBFBBFFRLR",
				Row: 44,
				Col: 5,
				ID:  357,
			},
		},
		{
			input: "BFFFBBFRRR",
			seat: Seat{
				Raw: "BFFFBBFRRR",
				Row: 70,
				Col: 7,
				ID:  567,
			},
		},
		{
			input: "FFFBBBFRRR",
			seat: Seat{
				Raw: "FFFBBBFRRR",
				Row: 14,
				Col: 7,
				ID:  119,
			},
		},
		{
			input: "BBFFBBFRLL",
			seat: Seat{
				Raw: "BBFFBBFRLL",
				Row: 102,
				Col: 4,
				ID:  820,
			},
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			seat := ConvertSeat(test.input)
			assert.Equal(t, test.seat.Raw, seat.Raw)
			assert.Equal(t, test.seat.Row, seat.Row)
			assert.Equal(t, test.seat.Col, seat.Col)
			assert.Equal(t, test.seat.ID, seat.ID)
		})
	}
}
