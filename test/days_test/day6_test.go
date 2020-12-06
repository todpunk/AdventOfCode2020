package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay6(t *testing.T) {
	var example = []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	var result int = days.Day6a(example)
	if result != 11 {
		t.Fail()
		t.Logf("Day6a example should return 11, got %d", result)
	}
	result = days.Day6b(example)
	if result != 6 {
		t.Fail()
		t.Logf("Day6b example should return 6, got %d", result)
	}
}
