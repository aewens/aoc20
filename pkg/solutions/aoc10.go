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

func (chain *AdapterChain) Link(adapter *Adapter) *AdapterChain {
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
		chain = chain.Link(adapter)
	}

	return diffs[1] * diffs[3]
}

func BuildChain(
		chains []*AdapterChain,
		chain *AdapterChain,
		values map[int]bool,
		builtin int,
	) []*AdapterChain {
	newChains := []*AdapterChain{}
	for c := 1; c <= 3; c++ {
		value := chain.Self.Jolts+c
		_, ok := values[value]
		if !ok {
			continue
		}

		newChain := &AdapterChain{
			Self: chain.Self,
			Prev: chain.Prev,
			Next: chain.Next,
		}
		newChain = newChain.Link(NewAdapter(value))
		if newChain.Self.Jolts == builtin {
			newChains = append(newChains, newChain)
		}

		chains = BuildChain(chains, newChain, values, builtin)
	}

	chains = append(chains, newChains...)
	return chains
}

func BuildAllChains(values map[int]bool) int {
	var emptyAdapter *Adapter = nil
	var emptyChain *AdapterChain = &AdapterChain{Self: emptyAdapter}

	keys := []int{}
	for key := range values {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Add our built-in adapter
	builtin := keys[len(keys)-1]+3
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

	chains := []*AdapterChain{}
	chains = BuildChain(chains, chain, values, builtin)

	//validChains := 0
	//for _, chain := range chains {
	//	if chain.Self.Jolts != builtin {
	//		continue
	//	}

	//	check := chain.Prev
	//	for {
	//		if check.Prev == nil {
	//			break
	//		}

	//		if check.Self.Jolts == 0 {
	//			validChains = validChains + 1
	//			break
	//		}

	//		check = check.Prev
	//	}
	//}

	return len(chains)
}

func Solution10(lines chan string) {
	values := make(map[int]bool)
	for line := range lines {
		value := shared.StringToInt(line)
		values[value] = true
	}

	diffCode := BuildLongestChain(values)
	Display(1, diffCode)

	validChains := BuildAllChains(values)
	Display(2, validChains)
}
