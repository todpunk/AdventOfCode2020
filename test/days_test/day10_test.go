package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay10(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day10a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day10a example should return 5, got %d", result)
	}
	result = days.Day10b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day10b example should return 5, got %d", result)
	}
}
