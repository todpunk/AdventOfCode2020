package days

import (
	"strconv"
	"strings"
)

func Day8a(input []string) int {
	var accumulator int64 = 0
	var current_op int64 = 0
	var looped = false
	cpy := make([]string, len(input))
	copy(cpy, input)
	for !looped {
		op := strings.Split(cpy[current_op], " ")
		cpy[current_op] = "dun 4"
		switch op[0] {
		case "nop":
			current_op++
		case "jmp":
			da_int, _ := strconv.ParseInt(op[1], 10, 64)
			current_op = current_op + da_int
		case "acc":
			da_int, _ := strconv.ParseInt(op[1], 10, 64)
			accumulator = accumulator + da_int
			current_op++
		case "dun":
			looped = true
		default:
			println("This should never occur, bad input")
		}
	}
	return int(accumulator)
}

func Day8b(input []string) int {
	var accumulator int64 = 0
	var current_op int64 = 0
	var finished = false
	var looped = false
	var is_potential = false
	var count = 0
	var potentials = 0
	var i int64 = 0
	for i = int64(len(input)) - 1; i >= 0 && !finished; i-- {
		cpy := make([]string, len(input))
		copy(cpy, input)
		potential_op := strings.Split(cpy[i], " ")
		da_int, _ := strconv.ParseInt(potential_op[1], 10, 64)
		switch potential_op[0] {
		case "jmp":
			if i+da_int < i {
				is_potential = true
			}
		case "nop":
			if i+da_int > i {
				is_potential = true
			}
		default:
			is_potential = false
		}
		if is_potential {
			potentials++
			if potential_op[0] == "nop" {
				cpy[i] = strings.Replace(cpy[i], "nop", "jmp", 1)
			} else {
				cpy[i] = strings.Replace(cpy[i], "jmp", "nop", 1)
			}
		}
		// Run the instructions if we have potential
		for !finished && !looped && is_potential {
			op := strings.Split(cpy[current_op], " ")
			cpy[current_op] = "dun 4"
			switch op[0] {
			case "nop":
				current_op++
			case "jmp":
				da_int, _ := strconv.ParseInt(op[1], 10, 64)
				current_op = current_op + da_int
			case "acc":
				da_int, _ := strconv.ParseInt(op[1], 10, 64)
				accumulator = accumulator + da_int
				current_op++
			case "dun":
				looped = true
			default:
				println("This should never occur, bad input", current_op, len(cpy))
				break
			}
			if int(current_op) == len(cpy) {
				finished = true
			}
		}
		if is_potential && !finished { // Reset everything that failed
			current_op = 0
			accumulator = 0
			is_potential = false
			looped = false
		}
		count++
	}
	println("looked at", count, "things,", potentials, "potentials with the instruction being", i+1, input[i+1])
	return int(accumulator)
}
