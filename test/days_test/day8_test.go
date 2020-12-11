package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay8(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day8a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day8a example should return 5, got %d", result)
	}
	result = days.Day8b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day8b example should return 5, got %d", result)
	}
}
