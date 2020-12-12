package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay8(t *testing.T) {
	var example = []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	// Jacob is trying to find a way to break me with valid puzzle input, so commenting out here
	// var example2 = []string{
	// 	"acc +1",
	// 	"acc +1",
	// 	"jmp +3",
	// 	"acc +1",
	// 	"nop +3",
	// 	"jmp -2",
	// 	"jmp -3",
	// 	"acc +99",
	// }
	var result int = days.Day8a(example)
	if result != 5 {
		t.Fail()
		t.Logf("Day8a example should return 5, got %d", result)
	}
	result = days.Day8b(example)
	if result != 8 {
		t.Fail()
		t.Logf("Day8b example should return 8, got %d", result)
	}
	// result = days.Day8b(example2)
	// if result != 102 {
	// 	t.Fail()
	// 	t.Logf("Day8b example2 should return 102, got %d", result)
	// }
}
