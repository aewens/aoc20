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
	referenceType
)

type EqFunc func(int, int) int
type EqToken struct {
	Type  int
	Value int
}
type EqStack struct {
	Tokens []*EqToken
	Size   int
}

type Equation struct {
	Result    int
	Level     int
	Pointer   int
	Tokens    *EqStack
	Stack     *EqStack
	Registers map[int]*EqStack
	Resolved  map[int]*EqToken
}

func NewEqStack() *EqStack {
	return &EqStack{
		Tokens: []*EqToken{},
		Size:   0,
	}
}

func NewEquation() *Equation {
	return &Equation{
		Result:    -1,
		Level:     0,
		Pointer:   0,
		Tokens:    NewEqStack(),
		Stack:     NewEqStack(),
		Registers: make(map[int]*EqStack),
		Resolved:  make(map[int]*EqToken),
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

		eq.Tokens.Push(token)
	}
}

func (stack *EqStack) Push(token *EqToken) {
	stack.Tokens = append(stack.Tokens, token)
	stack.Size = stack.Size + 1
}

func (stack *EqStack) PushLeft(token *EqToken) {
	stack.Tokens = append([]*EqToken{token}, stack.Tokens...)
	stack.Size = stack.Size + 1
}

// I know I should return an error here, but this code will not get used again
func (stack *EqStack) Pop() *EqToken {
	if stack.Size == 0 {
		panic("Stack is empty, cannot pop")
	}

	token := stack.Tokens[stack.Size-1]
	stack.Tokens = stack.Tokens[:stack.Size-1]
	stack.Size = stack.Size - 1
	return token
}

func (stack *EqStack) Yank() *EqToken {
	if stack.Size == 0 {
		panic("Stack is empty, cannot pop")
	}

	token := stack.Tokens[0]
	stack.Tokens = stack.Tokens[1:]
	stack.Size = stack.Size - 1
	return token
}

func (eq *Equation) Apply(stack *EqStack) int {
	ops := make(map[int]EqFunc)
	ops[sumType] = func(a int, b int) int {
		return a + b
	}
	ops[productType] = func(a int, b int) int {
		return a * b
	}

	value := -1
	var op EqFunc
	for stack.Size > 0 {
		token := stack.Yank()
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
	var stack *EqStack

	for eq.Stack.Size > 1 {
		token := eq.Stack.Pop()
		if token.Type == closeBracketType {
			stack = NewEqStack()
			continue
		}

		if token.Type != openBracketType {
			stack.PushLeft(token)
			continue
		}

		result = &EqToken{
			Type:  numericType,
			Value: eq.Apply(stack),
		}
		break
	}
	eq.Stack.Push(result)
}

func (eq *Equation) Exec(closing bool) {
	if eq.Level > 0 && !closing {
		return
	}
	if eq.Level == 0 && closing && eq.Stack.Size == 1 {
		_ = eq.Stack.Pop()
		return
	}
	if closing && eq.Stack.Size == 3 {
		_ = eq.Stack.Pop()
	}

	if eq.Stack.Size > 2 {
		eq.Build()
		if eq.Stack.Size > 2 {
			return
		}
	}

	numeric := eq.Stack.Pop()
	op := eq.Stack.Pop()

	switch op.Type {
	case sumType:
		eq.Result = eq.Result + numeric.Value
	case productType:
		eq.Result = eq.Result * numeric.Value
	}
}

func (eq *Equation) Step() {
	token := eq.Tokens.Yank()
	switch token.Type {
	case numericType:
		if eq.Result == -1 {
			eq.Result = token.Value
			return
		}

		eq.Stack.Push(token)
		if eq.Stack.Size == 1 {
			return
		}
		eq.Exec(false)
	case sumType:
		eq.Stack.Push(token)
	case productType:
		eq.Stack.Push(token)
	case openBracketType:
		if eq.Result == -1 {
			eq.Result = 0
			eq.Stack.Push(&EqToken{
				Type: sumType,
			})
		}
		eq.Level = eq.Level + 1
		eq.Stack.Push(token)
	case closeBracketType:
		eq.Level = eq.Level - 1
		eq.Stack.Push(token)
		eq.Exec(true)
	}
}

func (eq *Equation) Run1() {
	for eq.Tokens.Size > 0 {
		eq.Step()
	}
}

func (eq *Equation) Store(parent *EqStack, stack *EqStack) *EqStack {
	reference := &EqToken{
		Type:  referenceType,
		Value: eq.Pointer,
	}
	//Display(-50, reference)
	parent.Push(reference)

	eq.Registers[eq.Pointer] = stack
	eq.Pointer = eq.Pointer + 1
	return parent
}

func (eq *Equation) SubstituteBrackets(stack *EqStack) *EqStack {
	var substack *EqStack

	cache := NewEqStack()
	for stack.Size > 0 {
		token := stack.Pop()
		cache.PushLeft(token)

		if token.Type == closeBracketType {
			substack = NewEqStack()
			continue
		}

		if token.Type != openBracketType {
			substack.PushLeft(token)
			continue
		}

		if substack.Size > 0 {
			return eq.Store(stack, substack)
		}
	}

	return cache
}

func (eq *Equation) ReduceBrackets() {
	for eq.Tokens.Size > 0 {
		token := eq.Tokens.Yank()
		//Display(-10, token)
		switch token.Type {
		case numericType:
			eq.Stack.Push(token)
		case sumType:
			eq.Stack.Push(token)
		case productType:
			eq.Stack.Push(token)
		case openBracketType:
			eq.Level = eq.Level + 1
			eq.Stack.Push(token)
		case closeBracketType:
			eq.Level = eq.Level - 1
			eq.Stack.Push(token)
			eq.Stack = eq.SubstituteBrackets(eq.Stack)
		}
	}

	for eq.Stack.Size > 0 {
		token := eq.Stack.Yank()
		//Display(-11, token)
		eq.Tokens.Push(token)
	}
}

func (eq *Equation) ReduceSums(stack *EqStack) {
	for stack.Size > 0 {
		token := stack.Yank()
		//Display(-20, token)
		switch token.Type {
		case numericType:
			eq.Stack.Push(token)
		case sumType:
			prev := eq.Stack.Pop()
			next := stack.Yank()
			//Display(-21, next)
			if prev.Type == referenceType || next.Type == referenceType {
				substack := NewEqStack()
				substack.Push(prev)
				substack.Push(token)
				substack.Push(next)
				//Display(-22, substack.Size)
				eq.Stack = eq.Store(eq.Stack, substack)
			} else {
				eq.Stack.Push(&EqToken{
					Type:  numericType,
					Value: prev.Value + next.Value,
				})
			}
		case productType:
			eq.Stack.Push(token)
		case referenceType:
			eq.Stack.Push(token)
		}
	}

	for eq.Stack.Size > 0 {
		token := eq.Stack.Yank()
		//Display(-23, token)
		stack.Push(token)
	}
}

func (eq *Equation) ReduceProducts(stack *EqStack) int {
	result := -1
	for stack.Size > 0 {
		token := stack.Yank()
		//Display(-30, token)
		switch token.Type {
		case numericType:
			result = token.Value
		case sumType:
			next := stack.Yank()
			if next.Type == referenceType {
				next = eq.Resolve(next.Value)
			}
			//Display(-31, next.Value)
			result = result + next.Value
		case productType:
			next := stack.Yank()
			if next.Type == referenceType {
				next = eq.Resolve(next.Value)
			}
			//Display(-31, next.Value)
			result = result * next.Value
		case referenceType:
			result = eq.Resolve(token.Value).Value
		}
		//Display(-32, result)
	}
	return result
}

func (eq *Equation) Resolve(pointer int) *EqToken {
	result, ok := eq.Resolved[pointer]
	if ok {
		return result
	}
	register := eq.Registers[pointer]
	value := eq.ReduceProducts(register)
	result = &EqToken{
		Type:  numericType,
		Value: value,
	}
	eq.Resolved[pointer] = result
	return result
}

func (eq *Equation) Run2() {
	eq.ReduceBrackets()

	ps := eq.Pointer
	for p := 0; p < ps; p++ {
		register := eq.Registers[p]
		//Display(-p-100, register.Size)
		eq.ReduceSums(register)
	}

	//Display(-eq.Pointer-100, eq.Tokens.Size)
	eq.ReduceSums(eq.Tokens)

	ps = eq.Pointer
	for p := 0; p < ps; p++ {
		//Display(-p-200, eq.Registers[p].Size)
		_ = eq.Resolve(p)
		//Display(-p-300, eq.Resolved[p])
	}
	//Display(-400, eq.Tokens.Size)
	eq.Result = eq.ReduceProducts(eq.Tokens)
}

func SumEquations(equations []*Equation) int {
	result := 0
	for _, equation := range equations {
		result = result + equation.Result
	}
	return result
}

func Solution18(lines chan string) {
	equations1 := []*Equation{}
	equations2 := []*Equation{}
	for line := range lines {
		equation1 := NewEquation()
		equation1.Parse(line)
		equation1.Run1()

		equation2 := NewEquation()
		equation2.Parse(line)
		equation2.Run2()

		equations1 = append(equations1, equation1)
		equations2 = append(equations2, equation2)
	}

	Display(1, SumEquations(equations1))
	Display(2, SumEquations(equations2))
}
/*
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
$3 + 2 + 4 * 2
$3 = ($1 * $2 + 6)
$1 = (2 + 4 * 9)
$2 = (6 + 9 * 8 + 6)
*/
