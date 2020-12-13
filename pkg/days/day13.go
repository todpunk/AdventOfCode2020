package days

import (
	"strconv"
	"strings"
)

func Day13a(input []string) int {
	var earliest_bus_id int64
	earliest_bus_time := int64(0x7FFFFFFF)
	earliest_time, _ := strconv.ParseInt(input[0], 10, 64)
	for _, id := range strings.Split(input[1], ",") {
		if id == "x" {
			continue
		}
		this_id, _ := strconv.ParseInt(id, 10, 64)
		this_bus_earliest := int64(int64(earliest_time/this_id)+1) * this_id

		if this_bus_earliest < earliest_bus_time {
			earliest_bus_time = this_bus_earliest
			earliest_bus_id = this_id
		}
	}
	return int((earliest_bus_time - earliest_time) * earliest_bus_id)
}

func Day13b(input []string) int {
	// adapted from https://github.com/colinodell/advent-2020/blob/main/day13/day13.go
	// because I don't do mathy nonsense in puzzles like AoC
	// Parse the bus schedule.
	// A given bus will depart when the timestamp is evenly divisible by the bus ID, so we'll put those schedules
	// into a slice where the key is the timestamp offset and the value is the bus ID (factor to divide by).
	var buses []int
	for _, id := range strings.Split(input[1], ",") {
		if id == "x" {
			// Bus departure time doesn't matter, so we'll use a factor of 1 which will always match
			id = "1"
		}
		bus_id, _ := strconv.ParseInt(id, 10, 64)
		buses = append(buses, int(bus_id))
	}

	var timestamp = 1

	for {
		var timeToSkipIfNoMatch = 1
		valid := true

		for offset := 0; offset < len(buses); offset++ {
			// A given bus will depart when the timestamp is evenly divisible by the bus ID
			if (timestamp+offset)%buses[offset] != 0 {
				// No match here; abort and we'll try the next potential timestamp
				valid = false
				break
			}

			// This particular bus schedule matches, but we don't know if subsequent ones will. However, we
			// do know this when this particular schedule will match again, so let's keep track of that.
			// For example, if the first bus is Bus 7, we know it won't depart again for another 7 minutes,
			// so we'll skip ahead by 7 minutes and ignore the timestamps we know won't match.
			//
			// If we find a partial match where, say, 2 or 3 schedules match but not the whole thing, we
			// can calculate the next time those 2-3 schedules align by multiplying their values together;
			// worst case, we still have no match there, or best case we find yet another matching bus!
			//
			// For example, let's say we find a timestamp where the first two buses (7 and 11) align but none
			// of the others do; in that case, we know that buses 7 and 9 won't align again for another
			// 77 (7*11) minutes, so we'll skip ahead 77 minutes. Eventually we might find that now buses
			// 7, 11, and 13 align, but none others do. Well, that means the next time that these three buses
			// align will be in 7*11*13 minutes, so skip ahead by that much and try again there.
			//
			// This approach significantly speeds up the search and the speed improves the bigger your
			// partial match is!
			//
			// (Note that technically we'd need to find the LCM of those bus IDs, but luckily our inputs are
			// always prime numbers so the LCM will always equal the product of those IDs.)
			timeToSkipIfNoMatch *= buses[offset]
		}

		// Did we find a full match?
		if valid {
			return int(timestamp)
		}

		timestamp += timeToSkipIfNoMatch
	}
}
