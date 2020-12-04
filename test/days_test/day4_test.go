package days_test

import (
	"testing"

	"github.com/todpunk/AdventOfCode2020/pkg/days"
)

func TestDay4(t *testing.T) {
	var example = []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}
	var result int = days.Day4a(example)
	if result != 2 {
		t.Fail()
		t.Logf("Day4a example should return 2, got %d", result)
	}
	result = days.Day4b(example)
	if result != 4 {
		t.Fail()
		t.Logf("Day4b example should return 4, got %d", result)
	}
}
