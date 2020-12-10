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

func (chain *AdapterChain) Link(adapter *Adapter) *AdapterChain {
	result := &AdapterChain{
		Self: adapter,
		Prev: chain,
	}
	chain.Next = result

	return result
}

func BuildChain(values []int) (*AdapterChain, int) {
	var emptyAdapter *Adapter = nil
	var emptyChain *AdapterChain = &AdapterChain{Self: emptyAdapter}
	sort.Ints(values)

	// Add our built-in adapter
	values = append(values, values[len(values)-1]+3)

	diffs := make(map[int]int)
	diffs[1] = 0
	diffs[2] = 0
	diffs[3] = 0

	// Initial output is 0 jolts
	chain := &AdapterChain{
		Self: NewAdapter(0),
		Prev: emptyChain,
	}

	for _, value := range values {
		adapter := NewAdapter(value)
		diff, ok := adapter.Compat[chain.Self.Jolts]
		if !ok {
			return chain, -1
		}
		diffs[diff] = diffs[diff] + 1
		chain = chain.Link(adapter)
	}

	return chain, diffs[1] * diffs[3]
}

func Solution10(lines chan string) {
	values := []int{}
	for line := range lines {
		value := shared.StringToInt(line)
		values = append(values, value)
	}

	_, diffCode := BuildChain(values)
	Display(1, diffCode)
}
