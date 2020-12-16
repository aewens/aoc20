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
}

func NewTicketSystem() *TicketSystem {
	return &TicketSystem{
		Mode:   rulesMode,
		Rules:  make(map[string][]*TicketRule),
		Yours:  []int{},
		Others: [][]int{},
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
	invalids := []int{}
	for _, ticket := range ts.Others {
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
			}
		}
	}
	return Sum(invalids)
}

func Solution16(lines chan string) {
	ts := NewTicketSystem()
	for line := range lines {
		ts.Parse(line)
	}
	Display(1, ts.Check())
}
