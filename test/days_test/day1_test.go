package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay1(t *testing.T) {
	var result int64 = days.Day1()
	if result != 4 {
		t.Fail()
		t.Logf("Day1 doesn't return 4! OMGWTFBBQ!!!")
	}
}
