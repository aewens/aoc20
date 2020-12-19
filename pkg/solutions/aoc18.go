package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[18] = Solution18
}

const (
	numericType int = iota
	sumType
	productType
	openBracketType
	closeBracketType
)

type EqFunc func(int, int) int
type EqToken struct {
	Type  int
	Value int
}

type Equation struct {
	Result int
	Level  int
	Tokens []*EqToken
	Stack  []*EqToken
}

func NewEquation() *Equation {
	return &Equation{
		Result: -1,
		Level:  0,
		Tokens: []*EqToken{},
		Stack:  []*EqToken{},
	}
}

func (eq *Equation) Parse(line string) {
	for _, letter := range line {
		if letter == ' ' {
			continue
		}

		token := &EqToken{
			Value: -1,
		}
		switch letter {
		case '(':
			token.Type = openBracketType
		case ')':
			token.Type = closeBracketType
		case '+':
			token.Type = sumType
		case '*':
			token.Type = productType
		default:
			token.Type = numericType
			token.Value = shared.RuneToInt(letter)
		}

		//Display(-1, token)
		eq.Tokens = append(eq.Tokens, token)
	}
}

func (eq *Equation) Push(token *EqToken) {
	eq.Stack = append(eq.Stack, token)
}

// I know I should return an error here, but this code will not get used again
func (eq *Equation) Pop() *EqToken {
	if len(eq.Stack) == 0 {
		panic("Stack is empty, cannot pop")
	}

	size := len(eq.Stack)
	token := eq.Stack[size-1]
	eq.Stack = eq.Stack[:size-1]
	return token
}

func (eq *Equation) Yank() *EqToken {
	if len(eq.Stack) == 0 {
		panic("Stack is empty, cannot yank")
	}


	token := eq.Stack[0]
	eq.Stack = eq.Stack[1:]
	return token
}

func (eq *Equation) Apply(stack []*EqToken) int {
	ops := make(map[int]EqFunc)
	ops[sumType] = func(a int, b int) int {
		return a + b
	}
	ops[productType] = func(a int, b int) int {
		return a * b
	}

	value := -1
	var op EqFunc
	for len(stack) > 0 {
		token := stack[0]
		stack = stack[1:]
		switch token.Type {
		case numericType:
			if value == -1 {
				value = token.Value
				//Display(-10, token.Value)
				continue
			}
			//Display(-10, token.Value)
			value = op(value, token.Value)
			//Display(-20, value)
		case sumType:
			op = ops[sumType]
			//Display(-10, "+")
		case productType:
			op = ops[productType]
			//Display(-10, "*")
		}
	}
	return value
}

// 1 + (2 * 3) + (4 * (5 + 6))
// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
func (eq *Equation) Build() {
	var result *EqToken

	level := 0
	stacks := [][]*EqToken{}
	op := eq.Yank()
	Display(-31, op)
	for len(eq.Stack) > 0 {
		token := eq.Yank()
		Display(-31, token)
		if token.Type == openBracketType {
			level = level + 1
			stack := []*EqToken{}
			stacks = append(stacks, stack)
			continue
		}

		if level > 0 {
			size := len(stacks)
			stack := stacks[size-1]
			if token.Type != closeBracketType {
				stacks[size-1] = append(stack, token)
				continue
			}

			level = level - 1
			output := &EqToken{
				Type:  numericType,
				Value: eq.Apply(stack),
			}

			stacks = stacks[:size-1]
			size = len(stacks)
			Display(-4, []int{level, len(stacks), output.Type, output.Value})
			if level > 0 && size > 0 {
				stacks[size-1] = append(stacks[size-1], output)
			} else {
				result = output
				Display(-5, result)
			}
			continue
		}
	}
	if result == nil || len(stacks) > 0 {
		Display(-32, result)
		Display(-33, len(stacks))
		result = &EqToken{
			Type:  numericType,
			Value: eq.Apply(stacks[0]),
		}
	}
	Display(-6, result)
	eq.Stack = append(eq.Stack, op, result)
}

func (eq *Equation) Exec(closing bool) {
	if eq.Level > 0 && !closing {
		return
	}
	if eq.Level == 0 && closing && len(eq.Stack) == 1 {
		_ = eq.Pop()
		return
	}

	Display(-30, len(eq.Stack))
	if !closing && len(eq.Stack) > 2 || closing && len(eq.Stack) > 3 {
		eq.Build()
	}
	if closing && len(eq.Stack) == 3 {
		_ = eq.Pop()
	}
	Display(-34, len(eq.Stack))

	numeric := eq.Pop()
	op := eq.Pop()

	switch op.Type {
	case sumType:
		Display(-7, "+")
		eq.Result = eq.Result + numeric.Value
	case productType:
		Display(-8, "*")
		eq.Result = eq.Result * numeric.Value
	}
}

func (eq *Equation) Step() {
	token := eq.Tokens[0]
	eq.Tokens = eq.Tokens[1:]
	switch token.Type {
	case numericType:
		value := token.Value
		Display(-1, value)
		if eq.Result == -1 {
			eq.Result = value
			return
		}

		eq.Push(token)
		if len(eq.Stack) == 1 {
			return
		}
		eq.Exec(false)
	case sumType:
		Display(-1, "+")
		eq.Push(token)
	case productType:
		Display(-1, "*")
		eq.Push(token)
	case openBracketType:
		Display(-1, "(")
		if eq.Result == -1 {
			eq.Result = 0
			eq.Push(&EqToken{
				Type: sumType,
			})
		}
		eq.Level = eq.Level + 1
		eq.Push(token)
	case closeBracketType:
		Display(-1, ")")
		eq.Level = eq.Level - 1
		eq.Push(token)
		eq.Exec(true)
	}

	Display(-2, []int{eq.Result,len(eq.Stack)})
}

func (eq *Equation) Run() {
	for len(eq.Tokens) > 0 {
		eq.Step()
	}
	Display(-10, "----")
}

func SumEquations(equations []*Equation) int {
	result := 0
	for _, equation := range equations {
		result = result + equation.Result
	}
	return result
}

func Solution18(lines chan string) {
	equations := []*Equation{}
	for line := range lines {
		Display(0, line)
		equation := NewEquation()
		equation.Parse(line)
		equation.Run()

		equations = append(equations, equation)
	}

	Display(1, SumEquations(equations))
}

/*

((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2

(
(2+4*9)

(
54
*
(6+9*8+6)

(
54
*
126
+
6
)

6810
+
2

6812
+
4

6816
*
2

13632

*/
