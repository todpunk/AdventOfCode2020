package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay12(t *testing.T) {
	var example = []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	var result int = days.Day12a(example)
	if result != 25 {
		t.Fail()
		t.Logf("Day12a example should return 25, got %d", result)
	}
	result = days.Day12b(example)
	if result != 286 {
		t.Fail()
		t.Logf("Day12b example should return 286, got %d", result)
	}
}
