package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[9] = Solution9
}

func Sum(xs []int) int {
	y := 0
	for _, x := range xs {
		y = y + x
	}
	return y
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

func BreakCipher(values []int, check int) int {
	size := 2
	for {
		// Initialize buffer, 0 is used for dummy data
		buffer := []int{0}
		for s := 0; s < size-1; s++ {
			buffer = append(buffer, values[s])
		}

		for v := size-1; v < len(values); v++ {
			buffer = append(buffer[1:], values[v])
			if Sum(buffer) == check {
				min := -1
				max := -1
				for _, buf := range buffer {
					if min == -1 || buf < min {
						min = buf
					}

					if max == -1 || buf > max {
						max = buf
					}
				}
				return min + max
			}
		}

		size = size + 1
		if size == len(values) {
			break
		}
	}
	return -1
}

func Solution9(lines chan string) {
	result := -1
	values := []int{}
	preamble := []int{}
	for line := range lines {
		value := shared.StringToInt(line)
		values = append(values, value)

		buffer, invalid := ParseCipher(preamble, 25, value)
		preamble = buffer
		if invalid && result < 0 {
			Display(1, value)
			result = value
		}
	}

	weakness := BreakCipher(values, result)
	Display(2, weakness)
}
