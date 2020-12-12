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
	X         int
	Y         int
}

func (ferry *Ferry) Distance() int {
	x := ferry.X
	if x < 0 {
		x = -x
	}

	y := ferry.Y
	if y < 0 {
		y = -y
	}

	return x + y
}

func ParseCourse(ferry *Ferry, line string) {
	instruction := line[0]
	value := shared.StringToInt(line[1:])
	switch instruction {
	case 'N':
		ferry.Y = ferry.Y - value
	case 'S':
		ferry.Y = ferry.Y + value
	case 'W':
		ferry.X = ferry.X - value
	case 'E':
		ferry.X = ferry.X + value
	case 'L':
		shift := value / 90
		ferry.Direction = (ferry.Direction - shift) % 4
		if ferry.Direction < 0 {
			ferry.Direction = ferry.Direction + 4
		}
	case 'R':
		shift := value / 90
		ferry.Direction = (ferry.Direction + shift) % 4
	case 'F':
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
}

func Solution12(lines chan string) {
	ferry := &Ferry{Direction: 0, X: 0, Y: 0}
	for line := range lines {
		ParseCourse(ferry, line)
	}

	Display(1, ferry.Distance())
}
