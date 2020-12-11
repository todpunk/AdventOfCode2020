package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay11(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day11a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day11a example should return 5, got %d", result)
	}
	result = days.Day11b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day11b example should return 5, got %d", result)
	}
}
