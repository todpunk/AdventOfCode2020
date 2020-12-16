package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay17(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day17a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day17a example should return 5, got %d", result)
	}
	result = days.Day17b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day17b example should return 5, got %d", result)
	}
}
