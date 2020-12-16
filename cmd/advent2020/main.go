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
		fmt.Println("Day", *day, "Part 1:", days.Day1a(input, 2020))
		fmt.Println("Day", *day, "Part 2:", days.Day1b(input, 2020))
	case 2:
		var input = parse_from_file("./input_files/day2.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day2a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day2b(input))
	case 3:
		var input = parse_from_file("./input_files/day3.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day3a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day3b(input))
	case 4:
		var input = parse_from_file("./input_files/day4.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day4a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day4b(input))
	case 5:
		var input = parse_from_file("./input_files/day5.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day5a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day5b(input))
	case 6:
		var input = parse_from_file("./input_files/day6.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day6a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day6b(input))
	case 7:
		var input = parse_from_file("./input_files/day7.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day7a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day7b(input))
	case 8:
		var input = parse_from_file("./input_files/day8.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day8a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day8b(input))
	case 9:
		var input = convert_to_ints(parse_from_file("./input_files/day9.txt"))
		fmt.Println("Day", *day, "Part 1:", days.Day9a(input, 25))
		fmt.Println("Day", *day, "Part 2:", days.Day9b(input, 25))
	case 10:
		var input = convert_to_ints(parse_from_file("./input_files/day10.txt"))
		fmt.Println("Day", *day, "Part 1:", days.Day10a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day10b(input))
	case 11:
		var input = parse_from_file("./input_files/day11.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day11a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day11b(input))
	case 12:
		var input = parse_from_file("./input_files/day12.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day12a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day12b(input))
	case 13:
		var input = parse_from_file("./input_files/day13.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day13a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day13b(input))
	case 14:
		var input = parse_from_file("./input_files/day14.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day14a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day14b(input))
	case 15:
		var input = parse_from_file("./input_files/day15.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day15a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day15b(input))
	case 16:
		var input = parse_from_file("./input_files/day16.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day16a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day16b(input))
	case 17:
		var input = parse_from_file("./input_files/day17.txt")
		fmt.Println("Day", *day, "Part 1:", days.Day17a(input))
		fmt.Println("Day", *day, "Part 2:", days.Day17b(input))
	// case 18:
	// 	var input = parse_from_file("./input_files/day18.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day18a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day18b(input))
	// case 19:
	// 	var input = parse_from_file("./input_files/day19.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day19a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day19b(input))
	// case 20:
	// 	var input = parse_from_file("./input_files/day20.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day20a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day20b(input))
	// case 21:
	// 	var input = parse_from_file("./input_files/day21.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day21a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day21b(input))
	// case 22:
	// 	var input = parse_from_file("./input_files/day22.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day22a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day22b(input))
	// case 23:
	// 	var input = parse_from_file("./input_files/day23.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day23a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day23b(input))
	// case 24:
	// 	var input = parse_from_file("./input_files/day24.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day24a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day24b(input))
	// case 25:
	// 	var input = parse_from_file("./input_files/day25.txt")
	// 	fmt.Println("Day", *day, "Part 1:", days.Day25a(input))
	// 	fmt.Println("Day", *day, "Part 2:", days.Day25b(input))
	default:
		fmt.Println("We don't have that day...")
	}
}
