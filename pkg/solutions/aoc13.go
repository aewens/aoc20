package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[13] = Solution13
}

func EarliestBus(arrival string, buses string) int {
	target := shared.StringToInt(arrival)
	earliest := -1
	result := -1
	for _, bus := range Separate(buses, ",") {
		if bus == "x" {
			continue
		}
		id := shared.StringToInt(bus)
		div := target / id
		rem := target % id
		arrives := id * div
		if rem > 0 {
			arrives = arrives + id
		}
		if arrives >= target && (arrives < earliest || earliest == -1) {
			earliest = arrives
			result = id * (arrives - target)
		}
	}
	
	return result
}

func Solution13(lines chan string) {
	arrival := <-lines
	buses := <-lines
	earliest := EarliestBus(arrival, buses)
	Display(1, earliest)
}
