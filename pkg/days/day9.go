package days

func Day9a(input []int, preambleLength int) int {
	var found = 0
	for i := preambleLength; i < len(input) && found == 0; i++ {
		var target = input[i]
		var valid = false
		for j := i - preambleLength; j < i-1 && !valid; j++ {
			for k := j + 1; k < i && !valid; k++ {
				if input[j]+input[k] == target {
					valid = true
				}
			}
		}
		if !valid {
			found = input[i]
		}
	}
	return found
}

func Day9b(input []int) int {
	return 4
}
