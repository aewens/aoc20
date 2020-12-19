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
type EqStack []*EqToken

type Equation struct {
	Result int
	Level  int
	Tokens EqStack
	Stack  EqStack
}

func NewEquation() *Equation {
	return &Equation{
		Result: -1,
		Level:  0,
		Tokens: EqStack{},
		Stack:  EqStack{},
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

func (eq *Equation) Apply(stack EqStack) int {
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
				continue
			}
			value = op(value, token.Value)
		case sumType:
			op = ops[sumType]
		case productType:
			op = ops[productType]
		}
	}
	return value
}

func (eq *Equation) Build() {
	var result *EqToken
	var stack EqStack

	for len(eq.Stack) > 1 {
		token := eq.Pop()
		if token.Type == closeBracketType {
			stack = EqStack{}
			continue
		}

		if token.Type != openBracketType {
			stack = append(EqStack{token}, stack...)
			continue
		}

		result = &EqToken{
			Type:  numericType,
			Value: eq.Apply(stack),
		}
		break
	}
	eq.Stack = append(eq.Stack, result)
}

func (eq *Equation) Exec(closing bool) {
	if eq.Level > 0 && !closing {
		return
	}
	if eq.Level == 0 && closing && len(eq.Stack) == 1 {
		_ = eq.Pop()
		return
	}
	if closing && len(eq.Stack) == 3 {
		_ = eq.Pop()
	}

	if len(eq.Stack) > 2 {
		eq.Build()
		if len(eq.Stack) > 2 {
			return
		}
	}

	numeric := eq.Pop()
	op := eq.Pop()

	switch op.Type {
	case sumType:
		eq.Result = eq.Result + numeric.Value
	case productType:
		eq.Result = eq.Result * numeric.Value
	}
}

func (eq *Equation) Step() {
	token := eq.Tokens[0]
	eq.Tokens = eq.Tokens[1:]
	switch token.Type {
	case numericType:
		value := token.Value
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
		eq.Push(token)
	case productType:
		eq.Push(token)
	case openBracketType:
		if eq.Result == -1 {
			eq.Result = 0
			eq.Push(&EqToken{
				Type: sumType,
			})
		}
		eq.Level = eq.Level + 1
		eq.Push(token)
	case closeBracketType:
		eq.Level = eq.Level - 1
		eq.Push(token)
		eq.Exec(true)
	}
}

func (eq *Equation) Run() {
	for len(eq.Tokens) > 0 {
		eq.Step()
	}
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
		equation := NewEquation()
		equation.Parse(line)
		equation.Run()

		equations = append(equations, equation)
	}

	Display(1, SumEquations(equations))
}
