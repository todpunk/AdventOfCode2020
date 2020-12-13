package days

import (
	"math"
	"sort"
)

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func CalcAdapterCombinations(diffs []int) int64 {
	// Based off of JacobPuff's math which we made work with differences of 2
	// https://github.com/JacobPuff/AoC-2020-Rust/blob/master/AoC2020Rust/src/days/day10.rs
	var retVal int64 = 0
	if len(diffs) <= 2 {
		return 1
	}
	for i := 0; i < round(math.Pow(2, float64(len(diffs)-1))); i++ {
		var zeroed_total = 0
		var valid = true
		for idx := 0; idx < len(diffs)-1; idx++ {
			var bit = 1 << idx
			if i&bit == 0 {
				zeroed_total = zeroed_total + diffs[idx]
			} else {
				zeroed_total = 0
			}
			if zeroed_total+diffs[idx+1] > 3 {
				valid = false
				break
			}
		}
		if valid {
			retVal++
		}
	}
	return retVal
}

func Day10a(input []int) int64 {
	var ones, threes int64
	var prev = 0
	cpy := make([]int, len(input))
	copy(cpy, input)
	var full_input = append([]int{0}, cpy...)
	sort.Ints(full_input)
	for _, val := range full_input {
		if val-prev == 1 {
			ones++
		}
		if val-prev == 3 {
			threes++
		}
		prev = val
	}
	return ones * (threes + 1)
}

func Day10b(input []int) int64 {
	// var next_input, multiplier int
	var rolling_diffs = []int{}
	var diff_sections = [][]int{}
	var combinations int64 = 1
	var prev = 0
	cpy := make([]int, len(input))
	copy(cpy, input)
	var full_input = append([]int{0}, cpy...)
	sort.Ints(full_input)
	full_input = append(full_input, full_input[len(full_input)-1]+3)
	for _, val := range full_input {
		if val-prev == 3 && len(rolling_diffs) > 0 {
			diff_sections = append(diff_sections, rolling_diffs)
			rolling_diffs = []int{}
		} else {
			rolling_diffs = append(rolling_diffs, val-prev)
		}
		prev = val
	}
	for _, diffs := range diff_sections {
		combinations = combinations * CalcAdapterCombinations(diffs)
	}
	return combinations
}
