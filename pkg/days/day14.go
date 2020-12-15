package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ReplaceNth(s, old, new string, n int) string {
	// Yoinked wholesale: https://stackoverflow.com/questions/44144641/replace-nth-occurrence-of-string-in-golang
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

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
	var mem = map[uint64]uint64{}
	var mask = "000000000000000000000000000000000000"
	var mask_xs = 0
	var total uint64
	var bits = []uint64{1}
	for i := 0; i < 64; i++ {
		bits = append(bits, bits[i]*2)
	}
	for _, line := range input {
		if line[:3] == "mas" {
			mask = strings.Split(line, " = ")[1]
			mask_xs = 0
			for _, override := range mask {
				if override == 'X' {
					mask_xs++
				}
			}
		}
		if line[:3] == "mem" {
			address, _ := strconv.ParseUint(line[4:strings.Index(line, "]")], 10, 64)
			val, _ := strconv.ParseUint(strings.Split(line, " = ")[1], 10, 64)
			stradd := fmt.Sprintf("%036v", strconv.FormatUint(address, 2))

			for i, override := range mask {
				if override == '1' || override == 'X' {
					stradd = strings.Join([]string{stradd[:i], string(override), stradd[i+1:]}, "")
				}
			}
			// fmt.Println("Mask:", mask)
			// see day 10 for round implementation
			for count := round(math.Pow(2, float64(mask_xs)) - 1); count >= 0; count-- {
				next_address := stradd
				// make a new mask
				for i := 0; i < mask_xs; i++ {
					new_val := "0"
					// If the leftmost bit we're looking at in count is 1, set new_val
					if bits[mask_xs-i-1]&uint64(count) > 0 {
						new_val = "1"
					}
					// fmt.Println("Setx:", next_address, fmt.Sprintf("%05v", strconv.FormatUint(uint64(count), 2)), new_val, i+1)
					next_address = ReplaceNth(next_address, "X", new_val, 1)
				}
				// fmt.Println("Then:", next_address)
				this_address, _ := strconv.ParseUint(next_address, 2, 64)
				mem[this_address] = val
			}
		}
	}
	for _, val := range mem {
		total = total + val
	}
	return int(total)
}
