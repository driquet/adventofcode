package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	keys        = regexp.MustCompile(`(byr|iyr|eyr|hgt|hcl|ecl|pid):(\S*)`)
	yearRe      = regexp.MustCompile(`^\d{4}$`)
	heightRe    = regexp.MustCompile(`^(\d+)(cm|in)$`)
	hairColorRe = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eyesColors  = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	passportIDRe = regexp.MustCompile(`^\d{9}$`)
)

func validateField(key, value string) bool {
	switch key {
	case "byr":
		if !yearRe.MatchString(value) {
			return false
		}
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if year < 1920 || year > 2002 {
			return false
		}

	case "iyr":
		if !yearRe.MatchString(value) {
			return false
		}
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if year < 2010 || year > 2020 {
			return false
		}

	case "eyr":
		if !yearRe.MatchString(value) {
			return false
		}
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if year < 2020 || year > 2030 {
			return false
		}

	case "hgt":
		matches := heightRe.FindStringSubmatch(value)
		if len(matches) != 3 {
			return false
		}

		number, err := strconv.Atoi(matches[1])
		if err != nil {
			return false
		}
		unit := matches[2]

		switch unit {
		case "cm":
			if number < 150 || number > 193 {
				return false
			}

		case "in":
			if number < 59 || number > 76 {
				return false
			}
		}

	case "hcl":
		if !hairColorRe.MatchString(value) {
			return false
		}

	case "ecl":
		if !eyesColors[value] {
			return false
		}

	case "pid":
		if !passportIDRe.MatchString(value) {
			return false
		}
	}

	return true
}

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal("unable to read file")
	}

	// Separate passports data
	var passports []string
	var currentPassport string

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = ""
		}

		currentPassport += "\n" + line
	}

	fmt.Printf("got %d passports data\n", len(passports))

	// Part one, count passports with all required keys
	var count int
	for _, passport := range passports {
		fields := make(map[string]bool)
		matches := keys.FindAllStringSubmatch(passport, -1)
		for _, match := range matches {
			key := match[1]
			fields[key] = true
		}

		if len(fields) == 7 {
			count++
		}
	}
	fmt.Printf("part one: valid passports = %d\n", count)

	// Part two, count passports with all required keys, validation needed
	count = 0

	for _, passport := range passports {
		fields := make(map[string]bool)
		matches := keys.FindAllStringSubmatch(passport, -1)

		valid := true
		for _, match := range matches {
			key := match[1]
			value := match[2]

			if !validateField(key, value) {
				fmt.Printf("invalid field: %s=%s\n", key, value)
				valid = false
				break
			}

			fields[key] = true
		}

		if !valid {
			continue
		}

		if len(fields) == 7 {
			count++
		}
	}
	fmt.Printf("part two: valid passports = %d\n", count)
}
