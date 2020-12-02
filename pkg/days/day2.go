package days

import (
	"strconv"
	"strings"
)

func Day2a(input []string) int {
	var retval = 0
	var actual int64
	var splits []string
	var minmax_split []string
	var letter byte
	var password string
	for _, entry := range input {
		splits = strings.Split(entry, " ")
		minmax_split = strings.Split(splits[0], "-")
		letter = splits[1][0]
		password = splits[2]
		min, _ := strconv.ParseInt(minmax_split[0], 10, 64)
		max, _ := strconv.ParseInt(minmax_split[1], 10, 64)
		actual = 0
		for i := 0; i < len(password); i++ {
			if letter == password[i] {
				actual++
			}
		}
		if actual >= min && actual <= max {
			retval++
		}
	}
	return retval
}

func Day2b(input []string) int {
	var retval = 0
	var splits []string
	var leftright_split []string
	var letter byte
	var password string
	for _, entry := range input {
		splits = strings.Split(entry, " ")
		leftright_split = strings.Split(splits[0], "-")
		letter = splits[1][0]
		password = splits[2]
		left, _ := strconv.ParseInt(leftright_split[0], 10, 64)
		right, _ := strconv.ParseInt(leftright_split[1], 10, 64)
		if (letter == password[left-1]) != (letter == password[right-1]) {
			retval++
		}
	}
	return retval
}
