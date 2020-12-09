package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[9] = Solution9
}

func ParseCipher(preamble []int, size int, value int) ([]int, bool) {
	if len(preamble) < size {
		preamble = append(preamble, value)
		return preamble, false
	}

	invalid := CheckCipher(preamble, value)
	preamble = append(preamble[1:], value)
	return preamble, invalid
}

func CheckCipher(preamble []int, value int) bool {
	sums := make(map[int]bool)
	for n, number := range preamble {
		for nn, nNumber := range preamble {
			if n == nn {
				continue
			}

			sums[number+nNumber] = true
		}
	}

	_, ok := sums[value]
	return !ok
}

func Solution9(lines chan string) {
	preamble := []int{}
	for line := range lines {
		value := shared.StringToInt(line)
		buffer, invalid := ParseCipher(preamble, 25, value)
		preamble = buffer
		if invalid {
			Display(1, value)
			break
		}
	}
}
