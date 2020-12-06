package solutions

func init() {
	Map[6] = Solution6
}

func ParseResponses(lines []string) []int {
	seen := make(map[string]int)
	valid := make(map[string]bool)
	for _, line := range lines {
		for l := range line {
			letter := string(line[l])
			count, ok := seen[letter]
			if !ok {
				seen[letter] = 0
				count = 0
			}
			seen[letter] = count + 1
			if seen[letter] == len(lines) {
				valid[letter] = true
			}
		}
	}
	return []int{len(seen), len(valid)}
}

func Solution6(lines chan string) {
	groups := [][]string{}
	group := []string{}
	for line := range lines {
		if len(line) == 0 {
			groups = append(groups, group)
			group = []string{}
			continue
		}

		group = append(group, line)
	}

	count1 := 0
	count2 := 0
	for _, group := range groups {
		counts := ParseResponses(group)
		count1 = count1 + counts[0]
		count2 = count2 + counts[1]
	}

	Display(1, count1)
	Display(2, count2)
}
