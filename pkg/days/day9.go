package days

func getInvalidNumber(numlist []int, preambleLength int) (int, int) {
	var found = 0
	var i = 0
	for i = preambleLength; i < len(numlist) && found == 0; i++ {
		var target = numlist[i]
		var valid = false
		for j := i - preambleLength; j < i-1 && !valid; j++ {
			for k := j + 1; k < i && !valid; k++ {
				if numlist[j]+numlist[k] == target {
					valid = true
				}
			}
		}
		if !valid {
			found = numlist[i]
		}
	}
	return i, found
}

func Day9a(input []int, preambleLength int) int {
	_, invalidNum := getInvalidNumber(input, preambleLength)
	return invalidNum
}

func Day9b(input []int, preambleLength int) int {
	_, invalidNum := getInvalidNumber(input, preambleLength)
	var least = 4294967295
	var most = 0
	var total = 0
	for i := 0; i < len(input) && total != invalidNum; i++ {
		least = 4294967295
		most = 0
		total = 0
		for _, nextVal := range input[i:] {
			total = total + nextVal
			if total > invalidNum {
				break
			} else {
				if nextVal < least {
					least = nextVal
				}
				if nextVal > most {
					most = nextVal
				}
			}
			if total == invalidNum {
				break
			}
		}
	}
	return most + least
}
