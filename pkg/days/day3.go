package days

func find_trees(forest []string, right int, down int) int {
	var retval = 0
	var width = len(forest[0])
	var a int
	var spots = ".#"
	var column = 0
	for a = 0; a < len(forest); a = a + down {
		if forest[a][column%(width)] == spots[1] {
			retval++
		}
		column = column + right
	}
	return retval
}

func Day3a(input []string) int {
	return find_trees(input, 3, 1)
}

func Day3b(input []string) int {
	var a = find_trees(input, 1, 1)
	var b = find_trees(input, 3, 1)
	var c = find_trees(input, 5, 1)
	var d = find_trees(input, 7, 1)
	var e = find_trees(input, 1, 2)
	return a * b * c * d * e
}
