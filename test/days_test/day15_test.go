package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay15(t *testing.T) {
	var example = map[int][]string{
		436:  []string{"0,3,6"},
		1:    []string{"1,3,2"},
		10:   []string{"2,1,3"},
		27:   []string{"1,2,3"},
		78:   []string{"2,3,1"},
		438:  []string{"3,2,1"},
		1836: []string{"3,1,2"},
	}
	for k, v := range example {
		// if k == 436 {
		var result int = days.Day15a(v)
		if result != k {
			t.Fail()
			t.Logf("Day15a example should return %d, got %d", k, result)
		}
		// }
	}
	example = map[int][]string{
		175594:  []string{"0,3,6"},
		2578:    []string{"1,3,2"},
		3544142: []string{"2,1,3"},
		261214:  []string{"1,2,3"},
		6895259: []string{"2,3,1"},
		18:      []string{"3,2,1"},
		362:     []string{"3,1,2"},
	}
	for k, v := range example {
		result := days.Day15b(v)
		if result != k {
			t.Fail()
			t.Logf("Day15b example should return %d, got %d", k, result)
		}
	}
}
