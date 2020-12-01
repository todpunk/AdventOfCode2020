package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay1(t *testing.T) {
	var example = []int{1721, 979, 366, 299, 675, 1456}
	var sum_target = 2020
	var result int = days.Day1a(example, sum_target)
	if result != 514579 {
		t.Fail()
		t.Logf("Day1a example should return 514579, got %d", result)
	}
	result = days.Day1b(example, sum_target)
	if result != 241861950 {
		t.Fail()
		t.Logf("Day1b example should return 241861950, got %d", result)
	}
}
