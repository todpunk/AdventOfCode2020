package days

func Day6a(input []string) int {
	var answers = map[rune]bool{}
	var sum = 0
	for _, line := range input {
		if line == "" {
			sum = sum + len(answers)
			answers = map[rune]bool{}
		} else {
			for _, yes_question := range line {
				answers[yes_question] = true
			}
		}
	}
	sum = sum + len(answers)
	return sum
}

func Day6b(input []string) int {
	return 4
}
