package solutions

import (
	"strings"
	"regexp"
)

func init() {
	Map[19] = Solution19
}

const (
	literalType int = iota
	eitherType
	chainType
)

type MessageToken struct {
	Type int
	Value []string
}

type MessageSystem struct {
	Mode     int
	Chains   map[string][]string
	Tokens   map[string]*MessageToken
	Lookup   map[string]string
	Messages []string
}

func NewMessageSystem() *MessageSystem {
	return &MessageSystem{
		Mode:     0,
		Tokens:   make(map[string]*MessageToken),
		Lookup:   make(map[string]string),
		Messages: []string{},
	}
}

func (ms *MessageSystem) PushToken(line string) {
	if ms.Mode > 0 {
		ms.Messages = append(ms.Messages, line)
		return
	}

	if len(line) == 0 {
		ms.Mode = 1
		return
	}

	pair := Separate(line, ": ")
	index := pair[0]
	field := pair[1]
	value := []string{}
	var tokenType int
	if field[0] == '"' {
		tokenType = literalType
		literal := field[1:2]
		value = append(value, literal)
		//values = append(values, value)

		// Since this is a literal, we can already resolve it here
		ms.Lookup[index] = literal
	} else if strings.Contains(field, " | ") {
		side := Separate(field, " | ")

		left := side[0]
		ms.Tokens[index + "l"] = &MessageToken{
			Type: chainType,
			Value: Parameters(left),
		}

		right := side[1]
		ms.Tokens[index + "r"] = &MessageToken{
			Type: chainType,
			Value: Parameters(right),
		}

		tokenType = eitherType
		value = append(value, index + "l")
		value = append(value, index + "r")
	} else {
		tokenType = chainType
		value = Parameters(field)
	}

	ms.Tokens[index] = &MessageToken{
		Type: tokenType,
		Value: value,
	}
}

func (ms *MessageSystem) Expand(token *MessageToken) string {
	pattern := ""
	switch token.Type {
	case literalType:
		return token.Value[0]
	case eitherType:
		left := ms.Expand(ms.Tokens[token.Value[0]])
		right := ms.Expand(ms.Tokens[token.Value[1]])
		return "(" + left + "|" + right + ")"
	case chainType:
		for _, index := range token.Value {
			resolved := ms.Resolve(index)
			pattern = pattern + resolved
		}
		return pattern
	}
	return pattern
}

func (ms *MessageSystem) Resolve(index string) string {
	resolved, ok := ms.Lookup[index]
	if ok {
		return resolved
	}

	return ms.Expand(ms.Tokens[index])
}

func (ms *MessageSystem) Check() int {
	valid := 0
	pattern := ms.Resolve("0")
	regex := regexp.MustCompile("^" + pattern + "$")
	//Display(3, valids)
	for _, message := range ms.Messages {
		//Display(3, message)
		match := regex.Match([]byte(message))
		if match {
			valid = valid + 1
		}
	}
	return valid
}

func Solution19(lines chan string) {
	ms := NewMessageSystem()
	for line := range lines {
		ms.PushToken(line)
	}
	Display(1, ms.Check())
}
