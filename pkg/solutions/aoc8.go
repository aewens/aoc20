package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[8] = Solution8
}

const (
	nop int = iota
	acc
	jmp
)

type ConsoleOp struct {
	Type  int
	Value int
}

type Console struct {
	Accumulator int
	Position    int
	Memory      []*ConsoleOp
	Executed    map[int]bool
}

func NewConsole() *Console {
	return &Console{0, 0, []*ConsoleOp{}, make(map[int]bool)}
}

func (c *Console) Reset() {
	c.Accumulator = 0
	c.Position = 0
	c.Executed = make(map[int]bool)
}

func (c *Console) PushMemory(line string) {
	params := Parameters(line)

	var op int
	switch params[0] {
	case "nop":
		op = nop
	case "acc":
		op = acc
	case "jmp":
		op = jmp
	default:
		panic("Invalid operation")
	}

	sign := params[1][0]
	value := shared.StringToInt(params[1][1:])
	if sign == '-' {
		value = -value
	}

	c.Memory = append(c.Memory, &ConsoleOp{
		Type:  op,
		Value: value,
	})
}

func (c *Console) Step() bool {
	if c.Position == len(c.Memory) {
		return true
	}

	_, ok := c.Executed[c.Position]
	if ok {
		return true
	}
	c.Executed[c.Position] = true

	op := c.Memory[c.Position]
	switch op.Type {
	case nop:
		c.Position = c.Position + 1
	case acc:
		c.Accumulator = c.Accumulator + op.Value
		c.Position = c.Position + 1
	case jmp:
		c.Position = c.Position + op.Value
	}

	return false
}

func (c *Console) Run() int {
	for {
		if c.Step() {
			break
		}
	}
	return c.Accumulator
}

func (c *Console) Repair() int {
	for o, op := range c.Memory {
		if op.Type == acc {
			continue
		}
		if op.Type == nop {
			c.Memory[o] = &ConsoleOp{jmp, op.Value}
		} else {
			c.Memory[o] = &ConsoleOp{nop, op.Value}
		}

		c.Reset()
		value := c.Run()
		if c.Position == len(c.Memory) {
			return value
		}

		c.Memory[o] = op
	}

	return -1
}

func Solution8(lines chan string) {
	console := NewConsole()
	for line := range lines {
		console.PushMemory(line)
	}

	Display(1, console.Run())
	Display(2, console.Repair())
}
