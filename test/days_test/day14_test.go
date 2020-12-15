package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay14(t *testing.T) {
	var example = []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	var result int = days.Day14a(example)
	if result != 165 {
		t.Fail()
		t.Logf("Day14a example should return 165, got %d", result)
	}
	example = []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	result = days.Day14b(example)
	if result != 208 {
		t.Fail()
		t.Logf("Day14b example should return 208, got %d", result)
	}

	var replaceSource = "1X2X3X4X"
	var replaceText = "X"
	var newText = "YY"
	var expectedTexts = []string{"1YY2X3X4X", "1X2YY3X4X", "1X2X3YY4X", "1X2X3X4YY"}

	for i := 0; i < 4; i++ {
		resultText := days.ReplaceNth(replaceSource, replaceText, newText, i+1)
		if resultText != expectedTexts[i] {
			t.Fail()
			t.Logf("ReplaceNth should return %x, got %s", expectedTexts[i], resultText)
		}
	}
}
