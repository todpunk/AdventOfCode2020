package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay12(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day12a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day12a example should return 5, got %d", result)
	}
	result = days.Day12b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day12b example should return 5, got %d", result)
	}
}
