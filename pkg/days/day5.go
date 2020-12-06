package days

import (
	"strconv"
)

type boarding_pass struct {
	row    int64
	column int64
	id     int64
}

func parse_boarding_pass(pass string) (retPass boarding_pass) {
	var current_row = []rune(pass[:7])
	var current_column = []rune(pass[7:])
	for i, val := range current_row {
		if val == 'B' {
			current_row[i] = '1'
		} else {
			current_row[i] = '0'
		}
	}
	for i, val := range current_column {
		if val == 'R' {
			current_column[i] = '1'
		} else {
			current_column[i] = '0'
		}
	}
	retPass.row, _ = strconv.ParseInt(string(current_row), 2, 64)
	retPass.column, _ = strconv.ParseInt(string(current_column), 2, 64)
	retPass.id = (retPass.row * 8) + retPass.column
	return retPass
}

func Day5a(input []string) int {
	var largest_id int = 0
	var pass = boarding_pass{}
	for _, line := range input {
		pass = parse_boarding_pass(line)
		if int(pass.id) > largest_id {
			largest_id = int(pass.id)
		}
	}
	return largest_id
}

func Day5b(input []string) int {
	var existent_seats = map[int64]bool{}
	var my_id int64 = (128 * 8) + 7
	var pass = boarding_pass{}
	for _, line := range input {
		pass = parse_boarding_pass(line)
		if pass.id < my_id {
			my_id = pass.id
		}
		existent_seats[pass.id] = true
	}
	for found := false; !found || my_id >= (128*8)+7; {
		my_id++
		if _, ok := existent_seats[my_id]; !ok {
			if _, ok := existent_seats[my_id+1]; ok {
				found = true
			}
		}
	}
	return int(my_id)
}
