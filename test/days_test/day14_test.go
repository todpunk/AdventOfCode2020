package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay14(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day14a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day14a example should return 5, got %d", result)
	}
	result = days.Day14b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day14b example should return 5, got %d", result)
	}
}