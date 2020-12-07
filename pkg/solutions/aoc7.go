package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[7] = Solution7
}

type Bag struct {
	Count int
	Type  string
}

type Bags map[string][]*Bag

func ParseBags(line string, up Bags, down Bags) {
	// "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	rule := Separate(line[:len(line)-1], "s contain ")
	parent := &Bag{1, rule[0]}
	_, ok := down[parent.Type]
	if !ok {
		down[parent.Type] = []*Bag{}
	}

	children := Separate(rule[1], ", ")
	for _, child := range children {
		if child == "no other bags" {
			continue
		}

		params := Parameters(child)
		bag := &Bag{
			Count: shared.StringToInt(params[0]),
			Type:  Merge(params[1:]),
		}
		if bag.Type[len(bag.Type)-1] == 's' {
			bag.Type = bag.Type[:len(bag.Type)-1]
		}
		_, ok := up[bag.Type]
		if !ok {
			up[bag.Type] = []*Bag{}
		}

		up[bag.Type] = append(up[bag.Type], parent)
		down[parent.Type] = append(down[parent.Type], bag)
	}
}

func ValidParents(bags Bags, search string, seen map[string]bool) {
	parents, ok := bags[search]
	if !ok {
		return
	}

	for _, parent := range parents {
		_, ok := seen[parent.Type]
		if ok {
			continue
		}

		seen[parent.Type] = true
		ValidParents(bags, parent.Type, seen)
	}
}

func NeededBags(bags Bags, search string, needs int) int {
	children, ok := bags[search]
	if !ok {
		return needs
	}

	for _, child := range children {
		needs = needs + child.Count
		count := child.Count * NeededBags(bags, child.Type, 0)
		needs = needs + count
	}

	return needs
}

func Solution7(lines chan string) {
	up := make(Bags)
	down := make(Bags)
	for line := range lines {
		ParseBags(line, up, down)
	}

	search := "shiny gold bag"

	seen := make(map[string]bool)
	ValidParents(up, search, seen)
	Display(1, len(seen))

	needs := NeededBags(down, search, 0)
	Display(2, needs)
}
