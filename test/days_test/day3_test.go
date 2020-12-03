package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay3(t *testing.T) {
	var example = []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	var result int = days.Day3a(example)
	if result != 7 {
		t.Fail()
		t.Logf("Day3a example should return 7, got %d", result)
	}
	result = days.Day3b(example)
	if result != 336 {
		t.Fail()
		t.Logf("Day3b example should return 336, got %d", result)
	}
}
