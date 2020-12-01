package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay1(t *testing.T) {
	var result int = days.Day1([]int{1721, 979, 366, 299, 675, 1456})
	if result != 514579 {
		t.Fail()
		t.Logf("Day1 example should return 514579, got %d", result)
	}
}
