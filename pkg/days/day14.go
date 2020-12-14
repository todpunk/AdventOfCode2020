package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day14a(input []string) int {
	var mem = map[uint64]uint64{}
	var mask = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	var total uint64
	for _, line := range input {
		if line[:3] == "mas" {
			mask = strings.Split(line, " = ")[1]
		}
		if line[:3] == "mem" {
			address, _ := strconv.ParseUint(line[4:strings.Index(line, "]")], 10, 64)
			val, _ := strconv.ParseUint(strings.Split(line, " = ")[1], 10, 64)
			strval := fmt.Sprintf("%036v", strconv.FormatUint(val, 2))
			// fmt.Println(strval)
			for i, override := range mask {
				if override == '1' || override == '0' {
					strval = strings.Join([]string{strval[:i], string(override), strval[i+1:]}, "")
				}
			}
			// fmt.Println(mask)
			// fmt.Println(strval)
			mem[address], _ = strconv.ParseUint(strval, 2, 64)
			// fmt.Println(mem[address])
		}
	}
	for _, val := range mem {
		total = total + val
	}
	return int(total)
}

func Day14b(input []string) int {
	return 4
}
