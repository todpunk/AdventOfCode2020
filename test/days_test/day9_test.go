package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay9(t *testing.T) {
	var example = []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	var result int = days.Day9a(example, 5)
	if result != 127 {
		t.Fail()
		t.Logf("Day9a example should return 127, got %d", result)
	}
	result = days.Day9b(example, 5)
	if result != 62 {
		t.Fail()
		t.Logf("Day9b example should return 62, got %d", result)
	}
}
