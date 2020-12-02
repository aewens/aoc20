package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[2] = Solution2
}

type MinMax struct {
	Min int
	Max int
}

type PasswordPolicy struct {
	Check string
	Valid MinMax
}

type PasswordHistory struct {
	Password string
	Policy   PasswordPolicy
}

func ParseEntry(line string) PasswordHistory {
	rawHistory := Separate(line, ": ")
	rawPassword := rawHistory[1]

	rawPolicy := Separate(rawHistory[0], " ")
	rawValid := Separate(rawPolicy[0], "-")
	check := rawPolicy[1]

	policy := PasswordPolicy{
		Check: check,
		Valid: MinMax{
			Min: shared.StringToInt(rawValid[0]),
			Max: shared.StringToInt(rawValid[1]),
		},
	}
	return PasswordHistory{
		Password: rawPassword,
		Policy:   policy,
	}
}

func CheckPassword(history PasswordHistory) bool {
	seen := make(map[string]int)
	for _, character := range history.Password {
		letter := string(character)
		count, ok := seen[letter]
		if !ok {
			count = 0
		}
		seen[letter] = count + 1
	}

	found := false
	for letter, count := range seen {
		if letter != history.Policy.Check {
			continue
		}

		found = true
		if count < history.Policy.Valid.Min {
			return false
		}

		if count > history.Policy.Valid.Max {
			return false
		}
	}

	return found
}

func Solution2(lines chan string) {
	entries := []PasswordHistory{}
	for line := range lines {
		entry := ParseEntry(line)
		entries = append(entries, entry)
	}

	valid := 0
	for _, entry := range entries {
		if CheckPassword(entry) {
			valid = valid + 1
		}
	}

	Display(1, valid)
}
