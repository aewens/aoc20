package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[1] = Solution1
}

func Product(xs []int) int {
	result := 1
	for _, x := range xs {
		result = result * x
	}
	return result
}

func SearchDouble2020(transactions []int) []int {
	for t, transaction := range transactions {
		for tt, tTransaction := range transactions {
			if t == tt {
				continue
			}

			if transaction+tTransaction == 2020 {
				return []int{transaction, tTransaction}
			}
		}
	}
	return []int{-1, -1}
}

func SearchTriple2020(transactions []int) []int {
	for t, transaction := range transactions {
		for tt, tTransaction := range transactions {
			for ttt, ttTransaction := range transactions {
				if t == tt || t == ttt || tt == ttt {
					continue
				}

				if transaction+tTransaction+ttTransaction == 2020 {
					return []int{transaction, tTransaction, ttTransaction}
				}
			}
		}
	}
	return []int{-1, -1, -1}
}

func Solution1(lines chan string) {
	transactions := []int{}
	for line := range lines {
		transaction := shared.StringToInt(line)
		transactions = append(transactions, transaction)
	}

	Display(1, Product(SearchDouble2020(transactions)))
	Display(2, Product(SearchTriple2020(transactions)))
}
