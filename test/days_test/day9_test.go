package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay9(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day9a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day9a example should return 5, got %d", result)
	}
	result = days.Day9b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day9b example should return 5, got %d", result)
	}
}
