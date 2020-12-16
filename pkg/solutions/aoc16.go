package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[16] = Solution16
}

const (
	rulesMode int = iota
	yoursMode
	othersMode
)

type TicketRule struct {
	Lower int
	Upper int
}

type TicketSystem struct {
	Mode   int
	Rules  map[string][]*TicketRule
	Yours  []int
	Others [][]int
	Values [][]int
	Order  map[int]string
}

func All(xs []bool) bool {
	for _, x := range xs {
		if !x {
			return false
		}
	}
	return true
}

func NewTicketSystem() *TicketSystem {
	return &TicketSystem{
		Mode:   rulesMode,
		Rules:  make(map[string][]*TicketRule),
		Yours:  []int{},
		Others: [][]int{},
		Values: [][]int{},
		Order:  make(map[int]string),
	}
}

func (ts *TicketSystem) Parse(line string) {
	if len(line) == 0 {
		ts.Mode = ts.Mode + 1
		return
	}

	switch ts.Mode {
	case rulesMode:
		rules := []*TicketRule{}
		values := Separate(line, ": ")
		name := values[0]
		for _, conditions := range Separate(values[1], " or ") {
			bounds := Separate(conditions, "-")
			rule := &TicketRule{
				Lower: shared.StringToInt(bounds[0]),
				Upper: shared.StringToInt(bounds[1]),
			}
			rules = append(rules, rule)
		}
		ts.Rules[name] = rules
	case yoursMode:
		if line[len(line)-1] == ':' {
			return
		}

		yours := []int{}
		for _, value := range Separate(line, ",") {
			yours = append(yours, shared.StringToInt(value))
		}
		ts.Yours = yours
	case othersMode:
		if line[len(line)-1] == ':' {
			return
		}

		other := []int{}
		for _, value := range Separate(line, ",") {
			other = append(other, shared.StringToInt(value))
		}
		ts.Others = append(ts.Others, other)
	default:
		panic("Invalid mode")
	}
}

func (ts *TicketSystem) Check() int {
	keep := [][]int{}
	invalids := []int{}
	for _, ticket := range ts.Others {
		use := true
		for _, value := range ticket {
			valid := false
			for _, rules := range ts.Rules {
				for _, rule := range rules {
					if value >= rule.Lower && value <= rule.Upper {
						valid = true
						break
					}
				}
				if valid {
					break
				}
			}
			if !valid {
				invalids = append(invalids, value)
				use = false
			}
		}
		if use {
			keep = append(keep, ticket)
		}
	}
	ts.Others = keep
	return Sum(invalids)
}

func (ts *TicketSystem) SetValues() {
	// Initialize values
	for i := 0; i < len(ts.Others[0]); i++ {
		ts.Values = append(ts.Values, []int{})
	}

	// Populate values
	for _, ticket := range ts.Others {
		for v, value := range ticket {
			ts.Values[v] = append(ts.Values[v], value)
		}
	}
}

func (ts *TicketSystem) Match() {
	ts.SetValues()

	// Determine which rules are invalid for each index
	invalids := make(map[int]map[string]bool)
	for v, values := range ts.Values {
		_, ok := invalids[v]
		if !ok {
			invalids[v] = make(map[string]bool)
		}
		for _, value := range values {
			for name, rules := range ts.Rules {
				_, skip := invalids[v][name]
				if skip {
					continue
				}

				matched := false
				for _, rule := range rules {
					if value >= rule.Lower && value <= rule.Upper {
						matched = true
						break
					}
				}

				if !matched {
					invalids[v][name] = true
				}
			}
		}
	}

	// Keep looping to find which index only matches a single rule
	assigned := make(map[string]bool)
	for {
		if len(ts.Order) == len(ts.Rules) {
			break
		}

		valids := make(map[string][]int)
		for v := range ts.Values {
			_, set := ts.Order[v]
			if set {
				continue
			}

			for name := range ts.Rules {
				_, skip := assigned[name]
				if skip {
					continue
				}

				_, invalid := invalids[v][name]
				if invalid {
					continue
				}

				_, ok := valids[name]
				if !ok {
					valids[name] = []int{}
				}

				valids[name] = append(valids[name], v)
			}
		}

		for name, indices := range valids {
			if len(indices) == 1 {
				assigned[name] = true
				ts.Order[indices[0]] = name
			}
		}
	}
}

func (ts *TicketSystem) Departures() int {
	ts.Match()
	result := 1
	prefix := "departure"
	for index, name := range ts.Order {
		if len(name) < len(prefix) || name[:len(prefix)] != prefix {
			continue
		}
		value := ts.Yours[index]
		result = result * value
	}
	return result
}

func Solution16(lines chan string) {
	ts := NewTicketSystem()
	for line := range lines {
		ts.Parse(line)
	}
	Display(1, ts.Check())
	Display(2, ts.Departures())
}
