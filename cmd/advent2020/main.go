package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/todpunk/AdventOfCode2020/pkg/days"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	day = kingpin.Arg("day", "Advent day to run").Int()
)

func parse_from_file(filepath string) []string {
	var retData = []string{}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		retData = append(retData, scanner.Text())
	}
	return retData
}

func convert_to_ints(to_convert []string) []int {
	var retData = []int{}
	for i, _ := range to_convert {
		var val int64
		val, _ = strconv.ParseInt(to_convert[i], 10, 32)
		retData = append(retData, int(val))
	}
	return retData
}

func main() {
	kingpin.Version("0.1.0")
	kingpin.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Would run day: %d\n", *day)
	switch *day {
	case 1:
		var input = convert_to_ints(parse_from_file("./input_files/day1.txt"))
		fmt.Println("Day 1 Part 1:", days.Day1a(input, 2020))
		fmt.Println("Day 1 Part 2:", days.Day1b(input, 2020))
	case 2:
		var input = parse_from_file("./input_files/day2.txt")
		fmt.Println("Day 2 Part 1:", days.Day2a(input))
		fmt.Println("Day 2 Part 2:", days.Day2b(input))
	case 3:
		var input = parse_from_file("./input_files/day3.txt")
		fmt.Println("Day 3 Part 1:", days.Day3a(input))
		fmt.Println("Day 3 Part 2:", days.Day3b(input))
	case 4:
		var input = parse_from_file("./input_files/day4.txt")
		fmt.Println("Day 4 Part 1:", days.Day4a(input))
		fmt.Println("Day 4 Part 2:", days.Day4b(input))
	case 5:
		var input = parse_from_file("./input_files/day5.txt")
		fmt.Println("Day 5 Part 1:", days.Day5a(input))
		fmt.Println("Day 5 Part 2:", days.Day5b(input))
	case 6:
		var input = parse_from_file("./input_files/day6.txt")
		fmt.Println("Day 6 Part 1:", days.Day6a(input))
		fmt.Println("Day 6 Part 2:", days.Day6b(input))
	default:
		fmt.Println("We don't have that day...")
	}
}
