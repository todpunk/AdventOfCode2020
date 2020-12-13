package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func gibRunes(daStrings []string) []rune {
	var runes_list = []rune{}

	for _, line := range daStrings {
		runes_list = append(runes_list, []rune(line)...)
	}
	return runes_list
}

func TestDay11(t *testing.T) {
	var example = []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	var example2 = []string{
		"L.LL",
		"LLLL",
		"L.L.",
		"LLLL",
	}
	var result int = days.Day11a(example)
	if result != 37 {
		t.Fail()
		t.Logf("Day11a example should return 37, got %d", result)
	}
	result = days.Day11a(example2)
	if result != 6 {
		t.Fail()
		t.Logf("Day11a example should return 6, got %d", result)
	}
	result = days.Day11b(example)
	if result != 26 {
		t.Fail()
		t.Logf("Day11b example should return 26, got %d", result)
	}
	var seen_test1 = []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}
	var seen_test2 = []string{
		".............",
		".L.L.#.#.#.#.",
		".............",
	}
	var seen_test3 = []string{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"...L...",
		"##...##",
		"#.#.#.#",
		".##.##.",
	}
	result = days.OccupiedSeatsSeen(gibRunes(seen_test1), (4*len(seen_test1[0]))+3, len(seen_test1), len(seen_test1[0]))
	if result != 8 {
		t.Fail()
		t.Logf("OccupiedSeatsSeen seen_test1 should return 8, got %d", result)
	}
	result = days.OccupiedSeatsSeen(gibRunes(seen_test2), len(seen_test2[0])+1, len(seen_test2), len(seen_test2[0]))
	if result != 0 {
		t.Fail()
		t.Logf("OccupiedSeatsSeen seen_test2 should return 0, got %d", result)
	}
	result = days.OccupiedSeatsSeen(gibRunes(seen_test3), (3*len(seen_test3[0]))+3, len(seen_test3), len(seen_test3[0]))
	if result != 0 {
		t.Fail()
		t.Logf("OccupiedSeatsSeen seen_test3 should return 0, got %d", result)
	}
}
