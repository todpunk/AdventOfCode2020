package days

import (
	"fmt"
)

func seatsDifferent(a, b []rune) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}

func occupiedSeatsAdjacent(seats_list []rune, idx int, rows int, columns int) int {
	var occupied int = 0
	if idx >= columns && idx%columns != 0 {
		if seats_list[idx-columns-1] == '#' { // North-West occupied
			occupied++
		}
	}
	if idx >= columns {
		if seats_list[idx-columns] == '#' { // North occupied
			occupied++
		}
	}
	if idx >= columns && idx%columns != columns-1 {
		if seats_list[idx-columns+1] == '#' { // North-East occupied
			occupied++
		}
	}
	if idx%columns != 0 {
		if seats_list[idx-1] == '#' { // West occupied
			occupied++
		}
	}
	if idx%columns != columns-1 {
		if seats_list[idx+1] == '#' { // East occupied
			occupied++
		}
	}
	if idx < columns*(rows-1) && idx%columns != 0 {
		if seats_list[idx+columns-1] == '#' { // South-West occupied
			occupied++
		}
	}
	if idx < columns*(rows-1) {
		if seats_list[idx+columns] == '#' { // South occupied
			occupied++
		}
	}
	if idx < columns*(rows-1) && idx%columns != columns-1 {
		if seats_list[idx+columns+1] == '#' { // South-East occupied
			occupied++
		}
	}
	return occupied
}

func OccupiedSeatsSeen(seats_list []rune, idx int, rows int, columns int) int {
	var occupied int = 0

	for distance, pos := 0, idx; pos >= columns && pos%columns != 0; {
		distance++
		pos = idx - (columns * distance) - distance
		if seats_list[pos] == '#' { // North-West occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // North-West saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos >= columns; {
		distance++
		pos = idx - (columns * distance)
		if seats_list[pos] == '#' { // North occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos >= columns && pos%columns != columns-1; {
		distance++
		pos = idx - (columns * distance) + distance
		if seats_list[pos] == '#' { // North-East occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos%columns != 0; {
		distance++
		pos = idx - distance
		if seats_list[pos] == '#' { // West occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos%columns != columns-1; {
		distance++
		pos = idx + distance
		if seats_list[pos] == '#' { // East occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos < columns*(rows-1) && pos%columns != 0; {
		distance++
		pos = idx + (columns * distance) - distance
		if seats_list[pos] == '#' { // South-West occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos < columns*(rows-1); {
		distance++
		pos = idx + (columns * distance)
		if seats_list[pos] == '#' { // South occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	for distance, pos := 0, idx; pos < columns*(rows-1) && pos%columns != columns-1; {
		distance++
		pos = idx + (columns * distance) + distance
		if seats_list[pos] == '#' { // South-East occupied
			occupied++
			break
		}
		if seats_list[pos] == 'L' { // saw seat
			break
		}
	}
	return occupied
}

func Day11a(input []string) int {
	var columns = len(input[0])
	var rows = len(input)
	var seats_taken = 0
	var seats_list = []rune{}

	for _, line := range input {
		seats_list = append(seats_list, []rune(line)...)
	}
	var prev_seats_list = make([]rune, len(seats_list))
	var iter int
	for iter = 0; seatsDifferent(seats_list, prev_seats_list); iter++ {
		prev_seats_list = make([]rune, len(seats_list))
		copy(prev_seats_list, seats_list)
		for idx, spot := range seats_list {
			switch spot {
			case '.':
				continue
			case 'L': // Empty seat
				if occupiedSeatsAdjacent(prev_seats_list, idx, rows, columns) == 0 {
					seats_list[idx] = '#'
				}
			case '#': // Occupied seat
				if occupiedSeatsAdjacent(prev_seats_list, idx, rows, columns) >= 4 {
					seats_list[idx] = 'L'
				}
			default:
				fmt.Println("This should never be hit")
			}
		}
		// fmt.Println("damap:")
		// for i := 0; i < rows; i++ {
		// 	fmt.Println(string(seats_list[i*columns : ((i + 1) * columns)]))
		// }
	}

	// fmt.Println("Iterations:", iter)
	for _, seat := range seats_list {
		if seat == '#' {
			seats_taken++
		}
	}
	return seats_taken
}

func Day11b(input []string) int {
	var columns = len(input[0])
	var rows = len(input)
	var seats_taken = 0
	var seats_list = []rune{}

	for _, line := range input {
		seats_list = append(seats_list, []rune(line)...)
	}
	var prev_seats_list = make([]rune, len(seats_list))
	var iter int
	for iter = 0; seatsDifferent(seats_list, prev_seats_list); iter++ {
		prev_seats_list = make([]rune, len(seats_list))
		copy(prev_seats_list, seats_list)
		for idx, spot := range seats_list {
			switch spot {
			case '.':
				continue
			case 'L': // Empty seat
				if OccupiedSeatsSeen(prev_seats_list, idx, rows, columns) == 0 {
					seats_list[idx] = '#'
				}
			case '#': // Occupied seat
				if OccupiedSeatsSeen(prev_seats_list, idx, rows, columns) >= 5 {
					seats_list[idx] = 'L'
				}
			default:
				fmt.Println("This should never be hit")
			}
		}
		// fmt.Println("damap:")
		// for i := 0; i < rows; i++ {
		// 	fmt.Println(string(seats_list[i*columns : ((i + 1) * columns)]))
		// }
	}

	// fmt.Println("Iterations:", iter)
	for _, seat := range seats_list {
		if seat == '#' {
			seats_taken++
		}
	}
	return seats_taken
}
