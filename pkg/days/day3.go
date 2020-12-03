package days

func Day3a(input []string) int {
	var retval = 0
	var spots = ".#"
	var slope = 3
	var column = 0
	for _, forest := range input {
		if forest[column%(len(forest))] == spots[1] {
			retval++
		}
		column = column + slope
	}
	return retval
}

func Day3b(input []string) int {
	var retval = 7
	return retval
}
