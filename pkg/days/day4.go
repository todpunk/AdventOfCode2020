package days

import (
	"fmt"
	"strings"
)

func valid_passport(passport map[string]string) bool {
	var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var count = 0
	for _, field := range required {
		if _, ok := passport[field]; !ok {
			return false
		}
		count++
	}
	fmt.Printf("Valid Pasport found %d fields: %s\n", count, passport)
	return true
}

func parse_passport(passport string) map[string]string {
	var splits = strings.Split(passport, " ")
	var pairs = map[string]string{}
	for _, split := range splits {
		if len(split) > 1 {
			var pair = strings.Split(split, ":")
			pairs[pair[0]] = pair[1]
		}
	}
	return pairs
}

func Day4a(input []string) int {
	var valid_count = 0
	var current_passport = ""
	for i, line := range input {
		current_passport = current_passport + line + " "
		if line == "" || i+1 == len(input) { // Check the current passport as is, reset parser afterward
			if valid_passport(parse_passport(current_passport)) {
				valid_count++
			}
			current_passport = ""
		}
	}
	return valid_count
}

func Day4b(input []string) int {
	return 7
}
