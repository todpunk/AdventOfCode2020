package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/todpunk/AdventOfCode2020/pkg/days"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	day = kingpin.Arg("day", "Advent day to run").Int()
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Would run day: %d\n", *day)
	switch *day {
	case 1:
		fmt.Println("Day 1 Part 1:", days.Day1a(days.Day1_data, 2020))
		fmt.Println("Day 1 Part 2:", days.Day1b(days.Day1_data, 2020))
	case 2:
		fmt.Println("Day 2 Part 1:", days.Day2a(days.Day2_data))
		fmt.Println("Day 2 Part 2:", days.Day2b(days.Day2_data))
	case 3:
		fmt.Println("Day 3 Part 1:", days.Day3a(days.Day3_data))
		fmt.Println("Day 3 Part 2:", days.Day3b(days.Day3_data))
	case 4:
		fmt.Println("Day 4 Part 1:", days.Day4a(days.Day4_data))
		fmt.Println("Day 4 Part 2:", days.Day4b(days.Day4_data))
	case 5:
		fmt.Println("Day 5 Part 1:", days.Day5a(days.Day5_data))
		fmt.Println("Day 5 Part 2:", days.Day5b(days.Day5_data))
	case 6:
		fmt.Println("Day 6 Part 1:", days.Day6a(days.Day6_data))
		fmt.Println("Day 6 Part 2:", days.Day6b(days.Day6_data))
	default:
		fmt.Println("We don't have that day...")
	}
}
