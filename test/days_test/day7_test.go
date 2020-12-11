package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay7(t *testing.T) {
	var example = []string{
		"abc",
		"",
	}
	var result int = days.Day7a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day7a example should return 5, got %d", result)
	}
	result = days.Day7b(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day7b example should return 5, got %d", result)
	}
}
