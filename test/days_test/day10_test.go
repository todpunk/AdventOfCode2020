package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay10(t *testing.T) {
	var diffs = []int{1, 1, 1, 1}
	var comboResults = []int64{7, 13, 24, 44}
	for i := 0; i < 4; i++ {
		var diffRet = days.CalcAdapterCombinations(diffs)
		if diffRet != comboResults[i] {
			t.Fail()
			t.Logf("CalcAdapterCombinations diffs should return %d, got %d", comboResults[i], diffRet)
		}
		diffs = append(diffs, 1)
	}
	diffs = []int{1, 2, 1, 1}
	comboResults = []int64{5, 10, 18}
	for i := 0; i < 3; i++ {
		var diffRet = days.CalcAdapterCombinations(diffs)
		if diffRet != comboResults[i] {
			t.Fail()
			t.Logf("CalcAdapterCombinations diffs should return %d, got %d", comboResults[i], diffRet)
		}
		diffs = append(diffs, 1)
	}
	diffs = []int{1, 2, 2, 1, 1}
	comboResults = []int64{6, 12, 22}
	for i := 0; i < 3; i++ {
		var diffRet = days.CalcAdapterCombinations(diffs)
		if diffRet != comboResults[i] {
			t.Fail()
			t.Logf("CalcAdapterCombinations diffs should return %d, got %d", comboResults[i], diffRet)
		}
		diffs = append(diffs, 1)
	}
	diffs = []int{2, 2, 1, 1}
	comboResults = []int64{3, 6, 11}
	for i := 0; i < 3; i++ {
		var diffRet = days.CalcAdapterCombinations(diffs)
		if diffRet != comboResults[i] {
			t.Fail()
			t.Logf("CalcAdapterCombinations diffs should return %d, got %d", comboResults[i], diffRet)
		}
		diffs = append(diffs, 1)
	}
	var example = []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	var example2 = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	var result = days.Day10a(example)
	if result != 35 {
		t.Fail()
		t.Logf("Day10a example should return 35, got %d", result)
	}
	result = days.Day10a(example2)
	if result != 220 {
		t.Fail()
		t.Logf("Day10a example2 should return 5, got %d", result)
	}
	result = days.Day10b(example2)
	if result != 19208 {
		t.Fail()
		t.Logf("Day10b example should return 19208, got %d", result)
	}
}
