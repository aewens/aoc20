package solutions

import (
	"sort"

	"github.com/aewens/aoc20/pkg/shared"
)


func init() {
	Map[10] = Solution10
}

type Adapter struct {
	Jolts  int
	Compat map[int]int
}

type AdapterChain struct {
	Self *Adapter
	Prev *AdapterChain
	Next *AdapterChain
}

type AdapterChainMap map[string]*AdapterChain

func NewAdapter(jolts int) *Adapter {
	compat := make(map[int]int)
	compat[jolts-1] = 1
	compat[jolts-2] = 2
	compat[jolts-3] = 3
	return &Adapter{
		Jolts:  jolts,
		Compat: compat,
	}
}

func (chain *AdapterChain) Join(adapter *Adapter) *AdapterChain {
	result := &AdapterChain{
		Self: adapter,
		Prev: chain,
	}
	chain.Next = result

	return result
}

func BuildLongestChain(values map[int]bool) int {
	var emptyAdapter *Adapter = nil
	var emptyChain *AdapterChain = &AdapterChain{Self: emptyAdapter}

	keys := []int{}
	for key := range values {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Add our built-in adapter
	builtin := keys[len(keys)-1]+3
	keys = append(keys, builtin)
	values[builtin] = true

	diffs := make(map[int]int)
	diffs[1] = 0
	diffs[2] = 0
	diffs[3] = 0

	// Initial output is 0 jolts
	chain := &AdapterChain{
		Self: NewAdapter(0),
		Prev: emptyChain,
	}

	for _, value := range keys {
		adapter := NewAdapter(value)
		diff, ok := adapter.Compat[chain.Self.Jolts]
		if !ok {
			return -1
		}
		diffs[diff] = diffs[diff] + 1
		chain = chain.Join(adapter)
	}

	return diffs[1] * diffs[3]
}

func GenLinks(
		seen map[int]int,
		values map[int]bool,
		check int,
		stop int,
	) int {
	cache, skip := seen[check]
	if skip {
		return cache
	}

	counts := 0
	for c := 1; c <= 3; c++ {
		value := check-c
		_, ok := values[value]
		if !ok || value > stop {
			continue
		}

		count := GenLinks(seen, values, value, stop)
		counts = counts + count
	}
	seen[check] = counts
	return counts
}

func CountAllChains(values map[int]bool) int {
	keys := []int{}
	for key := range values {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Add our built-in adapter
	builtin := keys[len(keys)-1]+3
	keys = append(keys, builtin)
	values[builtin] = true
	values[0] = true

	seen := make(map[int]int)
	seen[0] = 1
	for _, key := range keys {
		GenLinks(seen, values, key, builtin)
	}

	return seen[builtin]
}

func Solution10(lines chan string) {
	values := make(map[int]bool)
	for line := range lines {
		value := shared.StringToInt(line)
		values[value] = true
	}

	diffCode := BuildLongestChain(values)
	Display(1, diffCode)

	validChains := CountAllChains(values)
	Display(2, validChains)
}
