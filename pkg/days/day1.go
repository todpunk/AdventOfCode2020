package days

func Day1a(input []int, target_sum int) int {
	var retval int
	var a, b int
	var found = false
	for a = 0; a < len(input)-1 && !found; a++ {
		for b = a + 1; b < len(input) && !found; b++ {
			if (input[a] + input[b]) == target_sum {
				found = true
			}
			if found {
				retval = input[a] * input[b]
			}
		}
	}
	return retval
}

func Day1b(input []int, target_sum int) int {
	var retval int
	var a, b, c int
	var found = false
	for a = 0; a < len(input)-2 && !found; a++ {
		for b = a + 1; b < len(input)-1 && !found; b++ {
			for c = b + 1; c < len(input) && !found; c++ {
				if (input[a] + input[b] + input[c]) == target_sum {
					found = true
				}
				if found {
					retval = input[a] * input[b] * input[c]
				}
			}
		}
	}
	return retval
}
