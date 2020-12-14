package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay15(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day15a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day15a example should return 5, got %d", result)
	}
	result = days.Day15b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day15b example should return 5, got %d", result)
	}
}
