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

func TestSolution3(t *testing.T) {
	lines := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	treeMap := &TreeMap{
		Pattern:  [][]bool{},
		Toboggan: &Point{0, 0},
		Vector: &Point{3, 1},
	}

	for _, line := range lines {
		ParsePattern(treeMap, line)
	}

	treeMap.Init()
	if len(lines[0]) != treeMap.Width {
		t.Fatalf("Day 3 parser is broken: %#v", treeMap)
	}

	if len(lines) != treeMap.Height {
		t.Fatalf("Day 3 parser is broken: %#v", treeMap)
	}

	expecting := []int{7, 336}
	treeMap.Descend()

	if treeMap.TreesHit != expecting[0] {
		t.Fatalf("Part 1 - Invalid count: %d", treeMap.TreesHit)
	}

	hits := treeMap.TreesHit
	vectors := []*Point{
		{1, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	shouldBe := []int{2, 3, 4, 2}
	for v, vector := range vectors {
		treeMap.Upgrade(vector)
		treeMap.Descend()
		hits = hits * treeMap.TreesHit

		if treeMap.TreesHit != shouldBe[v] {
			t.Fatalf("Part 2a - Invalid count: %d:%d", v, treeMap.TreesHit)
		}
	}

	if hits != expecting[1] {
		t.Fatalf("Part 2b - Invalid count: %d", hits)
	}
}

func TestSolution4(t *testing.T) {
	lines := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}
	passports := ParsePassports(lines)

	valid := 0
	expecting := 2
	for _, passport := range passports {
		if passport.Valid {
			valid = valid + 1
		}
	}

	if valid != expecting {
		t.Fatalf("Part 1 - Invalid count: %d", valid)
	}

	invalidLines := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	}
	invalidPassports := ParsePassports(invalidLines)

	valid1 := 0
	for _, invalidPassport := range invalidPassports {
		if !invalidPassport.Valid {
			continue
		}

		invalid := false
		for key, value := range invalidPassport.Fields {
			if !CheckPassportField(key, value) {
				invalid = true
				break
			}
		}

		if !invalid {
			valid1 = valid1 + 1
		}
	}

	if valid1 != 0 {
		t.Fatalf("Part 2a - Invalid count: %d", valid1)
	}

	validLines := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}
	validPassports := ParsePassports(validLines)

	valid2 := 0
	for _, validPassport := range validPassports {
		if !validPassport.Valid {
			continue
		}

		invalid := false
		for key, value := range validPassport.Fields {
			if !CheckPassportField(key, value) {
				Display(-1, key)
				Display(-2, value)
				invalid = true
				break
			}
		}

		if !invalid {
			valid2 = valid2 + 1
		}
	}

	if valid2 != len(validPassports) {
		t.Fatalf("Part 2b - Invalid count: %d", valid2)
	}
}

func TestSolution5(t *testing.T) {
	lines := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	expecting := []int{357, 567, 119, 820}

	partitions := []*Partition{}
	for _, line := range lines {
		partitions = append(partitions, ParsePartition(line))
	}

	for p, partition := range partitions {
		seat := SearchSeat(partition)
		if seat.Id != expecting[p] {
			t.Fatalf("Part 1 - Invalid seat - %#v", seat)
		}
	}
}

func TestSolution6(t *testing.T) {
	lines := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
		"",
	}

	groups := [][]string{}
	group := []string{}
	for _, line := range lines {
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

	expecting := []int{11, 6}
	if count1 != expecting[0] {
		t.Fatalf("Part 1 - Invalid count: %d", count1)
	}

	if count2 != expecting[1] {
		t.Fatalf("Part 2 - Invalid count: %d", count2)
	}
}

func TestSolution7(t *testing.T) {
	sets := [][]string{
		{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		},
		{
			
			"shiny gold bags contain 2 dark red bags.",
			"dark red bags contain 2 dark orange bags.",
			"dark orange bags contain 2 dark yellow bags.",
			"dark yellow bags contain 2 dark green bags.",
			"dark green bags contain 2 dark blue bags.",
			"dark blue bags contain 2 dark violet bags.",
			"dark violet bags contain no other bags.",
		},
	}

	setBagsUp := []Bags{}
	setBagsDown := []Bags{}
	for _, set := range sets {
		up := make(Bags)
		down := make(Bags)
		for _, line := range set {
			ParseBags(line, up, down)
		}
		setBagsUp = append(setBagsUp, up)
		setBagsDown = append(setBagsDown, down)
	}

	search := "shiny gold bag"

	seen := make(map[string]bool)
	ValidParents(setBagsUp[0], search, seen)
	count1 := len(seen)

	expecting1 := 4
	if count1 != expecting1 {
		t.Fatalf("Part 1 - Invalid count: %d", count1)
	}

	expecting2 := []int{32, 126}
	for s := range sets {
		count2 := NeededBags(setBagsDown[s], search, 0)

		if count2 != expecting2[s] {
			t.Fatalf("Part 2 - Invalid count: %d", count2)
		}
	}
}

func TestSolution8(t *testing.T) {
	lines := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	console := NewConsole()
	for _, line := range lines {
		console.PushMemory(line)
	}

	expecting := []int{5, 8}

	value1 := console.Run()
	if value1 != expecting[0] {
		t.Fatalf("Part 1 - Invalid count: %d", value1)
	}

	value2 := console.Repair()
	if value2 != expecting[1] {
		t.Fatalf("Part 2 - Invalid count: %d", value2)
	}
}
