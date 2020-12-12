package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[12] = Solution12
}

const (
	east int = iota
	south
	west
	north
)

type Ferry struct {
	Direction int
	Position  *Point
	Waypoint  *Point
}

type FerryAction struct {
	Op    int
	Value int
}

func (ferry *Ferry) Distance() int {
	x := ferry.Position.X
	if x < 0 {
		x = -x
	}

	y := ferry.Position.Y
	if y < 0 {
		y = -y
	}

	return x + y
}

func (ferry *Ferry) Process(action *FerryAction) {
	switch action.Op {
	case 0:
		ferry.Waypoint.Y = ferry.Waypoint.Y + action.Value
	case 1:
		ferry.Waypoint.Y = ferry.Waypoint.Y - action.Value
	case 2:
		ferry.Waypoint.X = ferry.Waypoint.X + action.Value
	case 3:
		ferry.Waypoint.X = ferry.Waypoint.X - action.Value
	case 4:
		shift := action.Value / 90
		wx := ferry.Waypoint.X
		wy := ferry.Waypoint.Y
		switch shift {
		case 1:
			ferry.Waypoint.X = -wy
			ferry.Waypoint.Y = wx
		case 2:
			ferry.Waypoint.X = -wx
			ferry.Waypoint.Y = -wy
		case 3:
			ferry.Waypoint.X = wy
			ferry.Waypoint.Y = -wx
		default:
			panic("Invalid rotation")
		}
	case 5:
		shift := action.Value / 90
		wx := ferry.Waypoint.X
		wy := ferry.Waypoint.Y
		switch shift {
		case 1:
			ferry.Waypoint.X = wy
			ferry.Waypoint.Y = -wx
		case 2:
			ferry.Waypoint.X = -wx
			ferry.Waypoint.Y = -wy
		case 3:
			ferry.Waypoint.X = -wy
			ferry.Waypoint.Y = wx
		default:
			panic("Invalid rotation")
		}
	case 6:
		for i := 0; i < action.Value; i++ {
			ferry.Position.X = ferry.Position.X + ferry.Waypoint.X
			ferry.Position.Y = ferry.Position.Y + ferry.Waypoint.Y
		}
	default:
		panic("Invalid action")
	}
}

func ParseCourse(ferry *Ferry, line string) *FerryAction {
	instruction := line[0]
	value := shared.StringToInt(line[1:])
	result := &FerryAction{
		Value: value,
	}
	switch instruction {
	case 'N':
		result.Op = 0
		ferry.Position.Y = ferry.Position.Y + value
	case 'S':
		result.Op = 1
		ferry.Position.Y = ferry.Position.Y - value
	case 'E':
		result.Op = 2
		ferry.Position.X = ferry.Position.X + value
	case 'W':
		result.Op = 3
		ferry.Position.X = ferry.Position.X - value
	case 'L':
		result.Op = 4
		shift := value / 90
		ferry.Direction = (ferry.Direction - shift) % 4
		if ferry.Direction < 0 {
			ferry.Direction = ferry.Direction + 4
		}
	case 'R':
		result.Op = 5
		shift := value / 90
		ferry.Direction = (ferry.Direction + shift) % 4
	case 'F':
		result.Op = 6

		// Marvel at just how lazy I am today
		switch ferry.Direction {
		case east:
			ParseCourse(ferry, "E"+line[1:])
		case south:
			ParseCourse(ferry, "S"+line[1:])
		case west:
			ParseCourse(ferry, "W"+line[1:])
		case north:
			ParseCourse(ferry, "N"+line[1:])
		}
	default:
		panic("Invalid instruction")
	}

	return result
}

func Solution12(lines chan string) {
	ferry := &Ferry{
		Direction: 0,
		Position:  &Point{0, 0},
		Waypoint:  &Point{10, 1},
	}
	actions := []*FerryAction{}
	for line := range lines {
		action := ParseCourse(ferry, line)
		actions = append(actions, action)
	}

	Display(1, ferry.Distance())

	ferry.Position.X = 0
	ferry.Position.Y = 0
	for _, action := range actions {
		ferry.Process(action)
	}

	Display(2, ferry.Distance())
}
