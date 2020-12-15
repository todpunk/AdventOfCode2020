package days

import (
	"strconv"
	"strings"
)

type day15Counter struct {
	turn  int
	count int
}

func play15Game(input string, nth_number int) int {
	var primers = map[int]*day15Counter{}
	var turn int = 1
	var last_spoken = 0
	for _, val := range strings.Split(input, ",") {
		num, _ := strconv.ParseInt(val, 10, 64)
		primers[int(num)] = &day15Counter{turn, 0}
		last_spoken = int(num)
		turn++
	}
	for ; turn <= nth_number; turn++ {
		prev_spokens_turn := primers[last_spoken].turn
		primers[last_spoken].turn = turn - 1
		if primers[last_spoken].count == 0 {
			// fmt.Println("Last Spoken was new:", last_spoken, "turn:", primers[last_spoken].turn, primers[last_spoken].count)
			primers[last_spoken].count++
			if _, ok := primers[0]; !ok {
				primers[0] = &day15Counter{turn, 0}
			} else {
				primers[0].count += 1
			}
			last_spoken = 0
			// fmt.Println("New Spoken is: 0")
		} else {
			// fmt.Println("Last Spoken isn't new:", last_spoken, "turn:", primers[last_spoken].turn, primers[last_spoken].count)
			primers[last_spoken].count++
			primers[last_spoken].turn = turn - 1
			new_spoken := turn - 1 - prev_spokens_turn
			// fmt.Println("New Spoken is:", new_spoken)
			if _, ok := primers[new_spoken]; !ok {
				primers[new_spoken] = &day15Counter{turn, 0}
			} else {
				primers[new_spoken].count += 1
			}
			last_spoken = new_spoken
		}
	}
	return last_spoken
}

func Day15a(input []string) int {
	return play15Game(input[0], 2020)
}

func Day15b(input []string) int {
	return play15Game(input[0], 30000000)
}
