package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay5(t *testing.T) {
	var example = []string{
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"FFFBBBFRLR",
		"FFFBBBFRLL",
		"FFFBBBFLRR",
		"FFFBBBFLRL",
		"FFFBBBFLLR",
		"FFFBBBFLLL",
		"BBFFBBFRLL",
	}
	var result int = days.Day5a(example)
	if result != 820 {
		t.Fail()
		t.Logf("Day5a example should return 820, got %d", result)
	}
	result = days.Day5b(example)
	if result != 118 {
		t.Fail()
		t.Logf("Day5b example should return 118, got %d", result)
	}
}
