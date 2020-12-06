package days

func everyone_answered(group []string, all bool) int {
	var sum = 0
	var group_answers = map[rune]int{}
	for _, answers := range group {
		for _, yes_question := range answers {
			if _, ok := group_answers[yes_question]; !ok {
				group_answers[yes_question] = 1
			} else {
				group_answers[yes_question] = group_answers[yes_question] + 1
			}
		}
	}
	for _, total := range group_answers {
		if all && total == len(group) {
			sum = sum + 1
		}
		if !all {
			sum = sum + 1
		}
	}
	return sum
}

func process_groups(groups []string, require_group_yes bool) int {
	var group = []string{}
	var sum = 0
	for _, line := range groups {
		if line != "" {
			group = append(group, line)
		} else {
			sum = sum + everyone_answered(group, require_group_yes)
			group = []string{}
		}
	}
	sum = sum + everyone_answered(group, require_group_yes)
	return sum
}

func Day6a(input []string) int {
	return process_groups(input, false)
}

func Day6b(input []string) int {
	return process_groups(input, true)
}
