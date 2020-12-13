package days

import (
	"fmt"
	"strconv"
)

type direction int
type turnDirection int

const (
	North direction = iota
	East
	South
	West
)

const (
	Right turnDirection = iota
	Left
)

type nautical_position struct {
	Dir direction
	x   int64
	y   int64
}

func absInt64(a int64) int64 {
	if a < 0 {
		return a * -1
	}
	return a
}

func turnShip(facing direction, dir turnDirection, degrees int64) direction {
	var lefts = map[int64]map[direction]direction{
		90:  map[direction]direction{North: West, East: North, South: East, West: South},
		180: map[direction]direction{North: South, East: West, South: North, West: East},
		270: map[direction]direction{North: East, East: South, South: West, West: North},
	}
	var rights = map[int64]map[direction]direction{
		90:  map[direction]direction{North: East, East: South, South: West, West: North},
		180: map[direction]direction{North: South, East: West, South: North, West: East},
		270: map[direction]direction{North: West, East: North, South: East, West: South},
	}
	var dataMap = lefts
	if dir == Right {
		dataMap = rights
	}
	switch degrees % 360 {
	case 0:
		return facing
	case 90:
		return dataMap[degrees][facing]
	case 180:
		return dataMap[degrees][facing]
	case 270:
		return dataMap[degrees][facing]
	default:
		fmt.Println("This should not be reachable with a valid dataset")
		return North
	}
}

func (w *nautical_position) turnWaypoint(dir turnDirection, degrees int64) {
	var temp int64 = w.x
	switch degrees % 360 {
	case 0:
		return
	case 90:
		if dir == Right {
			w.x = w.y
			w.y = temp * -1
		} else { // Left
			w.x = w.y * -1
			w.y = temp
		}
	case 180:
		w.x = w.x * -1
		w.y = w.y * -1
	case 270:
		if dir == Right {
			w.x = w.y * -1
			w.y = temp
		} else { // Left
			w.x = w.y
			w.y = temp * -1
		}
	default:
		fmt.Println("This should not be reachable with a valid dataset")
	}
}

func (p *nautical_position) movePosition(dir direction, distance int64) {
	switch dir {
	case North:
		p.y = p.y + distance
	case South:
		p.y = p.y - distance
	case East:
		p.x = p.x + distance
	case West:
		p.x = p.x - distance
	}
}

func (p *nautical_position) moveWaypoint(waypoint nautical_position, times int64) {
	p.y = p.y + (waypoint.y * times)
	p.x = p.x + (waypoint.x * times)
}

func Day12a(input []string) int {
	var ship = nautical_position{East, 0, 0}

	for _, instruction := range input {
		letter := instruction[:1]
		num, _ := strconv.ParseInt(instruction[1:], 10, 64)
		switch letter {
		case "N":
			ship.movePosition(North, num)
		case "E":
			ship.movePosition(East, num)
		case "S":
			ship.movePosition(South, num)
		case "W":
			ship.movePosition(West, num)
		case "F":
			ship.movePosition(ship.Dir, num)
		case "L":
			ship.Dir = turnShip(ship.Dir, Left, num)
		case "R":
			ship.Dir = turnShip(ship.Dir, Right, num)
		default:
			fmt.Println("This letter should not be in a valid dataset")
		}
	}
	return int(absInt64(ship.x) + absInt64(ship.y))
}

func Day12b(input []string) int {
	var ship = nautical_position{East, 0, 0}
	var waypoint = nautical_position{East, 10, 1}

	for _, instruction := range input {
		letter := instruction[:1]
		num, _ := strconv.ParseInt(instruction[1:], 10, 64)
		switch letter {
		case "N":
			waypoint.movePosition(North, num)
		case "E":
			waypoint.movePosition(East, num)
		case "S":
			waypoint.movePosition(South, num)
		case "W":
			waypoint.movePosition(West, num)
		case "F":
			ship.moveWaypoint(waypoint, num)
		case "L":
			waypoint.turnWaypoint(Left, num)
		case "R":
			waypoint.turnWaypoint(Right, num)
		default:
			fmt.Println("This letter should not be in a valid dataset")
		}
	}
	return int(absInt64(ship.x) + absInt64(ship.y))
}
