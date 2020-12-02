package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay2(t *testing.T) {
	var example = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	var result int = days.Day2a(example)
	if result != 2 {
		t.Fail()
		t.Logf("Day1a example should return 2, got %d", result)
	}
	result = days.Day2b(example)
	if result != 2 {
		t.Fail()
		t.Logf("Day1b example should return 2, got %d", result)
	}
}
