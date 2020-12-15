package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay16(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day16a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day16a example should return 5, got %d", result)
	}
	result = days.Day16b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day16b example should return 5, got %d", result)
	}
}
