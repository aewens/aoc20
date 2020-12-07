package solutions

func init() {
	Map[7] = Solution7
}

type Bags map[string][]string

func ParseBags(bags Bags, line string) {
	// "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	rule := Separate(line[:len(line)-1], "s contain ")
	parent := rule[0]
	children := Separate(rule[1], ", ")
	for _, child := range children {
		if child == "no other bags" {
			continue
		}

		params := Parameters(child)
		//bagCount := params[0]
		bagType := Merge(params[1:])
		if bagType[len(bagType)-1] == 's' {
			bagType = bagType[:len(bagType)-1]
		}
		_, ok := bags[bagType]
		if !ok {
			bags[bagType] = []string{}
		}

		bags[bagType] = append(bags[bagType], parent)
	}
}

func ValidParents(bags Bags, search string, seen map[string]bool) {
	parents, ok := bags[search]
	if !ok {
		return
	}

	for _, parent := range parents {
		_, ok := seen[parent]
		if ok {
			continue
		}

		seen[parent] = true
		ValidParents(bags, parent, seen)
	}
}

func Solution7(lines chan string) {
	bags := make(Bags)
	for line := range lines {
		ParseBags(bags, line)
	}

	seen := make(map[string]bool)
	ValidParents(bags, "shiny gold bag", seen)
	Display(1, len(seen))
}

