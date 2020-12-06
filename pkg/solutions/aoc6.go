package solutions

func init() {
	Map[6] = Solution6
}

func ParseResponses(lines []string) int {
	seen := make(map[string]bool)
	for _, line := range lines {
		for l := range line {
			letter := string(line[l])
			_, ok := seen[letter]
			if !ok {
				seen[letter] = true
			}
		}
	}
	return len(seen)
}

func Solution6(lines chan string) {
	groups := [][]string{}
	group := []string{}
	for line := range lines {
		if len(line) == 0 {
			groups = append(groups, group)
			group = []string{}
		}

		group = append(group, line)
	}

	count := 0
	for _, group := range groups {
		count = count + ParseResponses(group)
	}

	Display(1, count)
}
