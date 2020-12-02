package solutions

import (
	"testing"
)

//func Cleanup(t *testing.T) {
//	r := recover()
//	if r != nil {
//		t.Fatal(r)
//	}
//}

func TestSolution1(t *testing.T) {
	transactions := []int{1721, 979, 366, 299, 675, 1456}
	
	part1 := Product(SearchDouble2020(transactions))
	if part1 != 514579 {
		t.Fatal("Double search got wrong answer")
	}

	part2 := Product(SearchTriple2020(transactions))
	if part2 != 241861950 {
		t.Fatal("Triple search got wrong answer")
	}
}

func TestSolution2(t *testing.T) {
	lines := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	entries := []PasswordHistory{}
	for _, line := range lines {
		entry := ParseEntry(line)
		entries = append(entries, entry)
	}

	valid1 := 0
	valid2 := 0
	expecting1 := 2
	expecting2 := 1
	for _, entry := range entries {
		if CheckPassword1(entry) {
			valid1 = valid1 + 1
		}

		if CheckPassword2(entry) {
			valid2 = valid2 + 1
		}
	}

	if valid1 != expecting1 {
		t.Fatalf("Part 1 - Invalid count: %d", valid1)
	}

	if valid2 != expecting2 {
		t.Fatalf("Part 2 - Invalid count: %d", valid1)
	}
}
