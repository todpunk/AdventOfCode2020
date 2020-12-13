package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay13(t *testing.T) {
	var example = []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	var result int = days.Day13a(example)
	if result != 295 {
		t.Fail()
		t.Logf("Day13a example should return 295, got %d", result)
	}
	result = days.Day13b(example)
	if result != 1068781 {
		t.Fail()
		t.Logf("Day13b example should return 1068781, got %d", result)
	}
}
