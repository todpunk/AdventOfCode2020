package days

import (
	"strconv"
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

func v_birth(birth_year string) bool {
	if len(birth_year) != 4 {
		return false
	}
	var year_int, err = strconv.ParseInt(birth_year, 10, 64)
	if err != nil {
		return false
	}
	if year_int < 1920 || year_int > 2002 {
		return false
	}
	return true
}

func v_issue(issue_year string) bool {
	if len(issue_year) != 4 {
		return false
	}
	var year_int, err = strconv.ParseInt(issue_year, 10, 64)
	if err != nil {
		return false
	}
	if year_int < 2010 || year_int > 2020 {
		return false
	}
	return true
}

func v_expiration(expiration string) bool {
	if len(expiration) != 4 {
		return false
	}
	var year_int, err = strconv.ParseInt(expiration, 10, 64)
	if err != nil {
		return false
	}
	if year_int < 2020 || year_int > 2030 {
		return false
	}
	return true
}

func v_height(height string) bool {
	var in_cm = height[len(height)-2:]
	var units, err = strconv.ParseInt(height[:len(height)-2], 10, 64)
	if err != nil {
		return false
	}
	switch in_cm {
	case "in":
		if units < 59 || units > 76 {
			return false
		}
	case "cm":
		if units < 150 || units > 193 {
			return false
		}
	default:
		return false
	}
	return true
}

func v_hair_color(hair_color string) bool {
	if hair_color[0] != '#' {
		return false
	}
	var valid_chars = map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'a': true,
		'b': true,
		'c': true,
		'd': true,
		'e': true,
		'f': true,
	}
	var hex_hair = hair_color[1:]
	for _, this_char := range hex_hair {
		if _, ok := valid_chars[this_char]; !ok {
			return false
		}
	}
	return true
}

func v_eye_color(eye_color string) bool {
	var valid_colors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, valid_color := range valid_colors {
		if valid_color == eye_color {
			return true
		}
	}
	return false
}

func v_passport_id(passport_id string) bool {
	if len(passport_id) != 9 {
		return false
	}
	var _, err = strconv.ParseInt(passport_id, 10, 64)
	if err != nil {
		return false
	}
	return true
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
	var valid_count = 0
	var current_passport = ""
	for i, line := range input {
		current_passport = current_passport + line + " "
		if line == "" || i+1 == len(input) { // Check the current passport as is, reset parser afterward
			var passport = parse_passport(current_passport)
			if valid_passport(passport) {
				// We can validate fields now
				if v_birth(passport["byr"]) &&
					v_issue(passport["iyr"]) &&
					v_expiration(passport["eyr"]) &&
					v_height(passport["hgt"]) &&
					v_hair_color(passport["hcl"]) &&
					v_eye_color(passport["ecl"]) &&
					v_passport_id(passport["pid"]) {
					valid_count++
				}
			}
			current_passport = ""
		}
	}
	return valid_count
}
