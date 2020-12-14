package solutions

import (
	"math"

	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[14] = Solution14
}

// pun intended
type Docker struct {
	Bitmask [36]int
	Memory  map[int]int
}

func exp(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func NewDocker() *Docker {
	return &Docker{
		Memory:  make(map[int]int),
	}
}

func (docker *Docker) Read(line string) {
	params := Parameters(line)
	if len(params) != 3 {
		panic("Invalid input")
	}

	op := params[0]
	value := params[2]

	if op[len(op)-1] == ']' {
		// mem
		address := shared.StringToInt(op[4:len(op)-1])
		input := shared.StringToInt(value)

		encode := [36]int{}
		for e := 35; e >= 0; e-- {
			check := exp(2, e)
			if input < check {
				encode[e] = 0
				continue
			}

			encode[e] = 1
			input = input - check
		}

		result := 0
		for index := range encode {
			power := exp(2, index)
			mask := docker.Bitmask[index]
			if mask != 2 {
				encode[index] = mask
			}

			if encode[index] == 1 {
				result = result + power
			}
		}
		docker.Memory[address] = result
	} else {
		// mask
		docker.Bitmask = [36]int{}
		bits := value
		for v := 0; v < len(value); v++ {
			bit := bits[len(bits)-1:]
			bits = bits[:len(bits)-1]
			if bit == "X" {
				docker.Bitmask[v] = 2
				continue
			}
			docker.Bitmask[v] = shared.StringToInt(bit)
		}
	}
}

func (docker *Docker) Sum() int {
	result := 0
	for _, value := range docker.Memory {
		result = result + value
	}
	return result
}

func Solution14(lines chan string) {
	docker := NewDocker()
	for line := range lines {
		docker.Read(line)
	}

	Display(1, docker.Sum())
}
