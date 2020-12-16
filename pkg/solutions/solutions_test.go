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

func TestSolution9(t *testing.T) {
	values := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	expecting := []int{127, 62}

	result := -1
	preamble := []int{}
	for _, value := range values {
		buffer, invalid := ParseCipher(preamble, 5, value)
		preamble = buffer
		if invalid {
			if value != expecting[0] {
				t.Fatalf("Part 1a - Invalid result: %d", value)
			}
			result = value
		}
	}

	if result == -1 {
		t.Fatalf("Part 1b - Did not find invalid")
	}

	weakness := BreakCipher(values, result)
	if weakness != expecting[1] {
		t.Fatalf("Part 2 - Invalid result: %d", weakness)
	}
}

func TestSolution10(t *testing.T) {
	sets := [][]int{
		{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		},
		{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		},
	}

	expecting1 := []int{7*5, 22*10}
	expecting2 := []int{8, 19208}

	for s, set := range sets {
		values := make(map[int]bool)
		for _, value := range set {
			values[value] = true
		}
		diffCode := BuildLongestChain(values)
		if diffCode != expecting1[s] {
			t.Fatalf("Part 1 - Invalid value: %d", diffCode)
		}

		validChains := CountAllChains(values)
		if validChains != expecting2[s] {
			t.Fatalf("Part 2 - Invalid value: %d", validChains)
		}
	}
}

func TestSolution11(t *testing.T) {
	lines := []string {
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	game := &GameOfSeats{
		Seats: make(map[int]map[int]int),
		Copy: make(map[int]map[int]int),
	}
	for _, line := range lines {
		ParseGame(game, line)
	}
	game.Height = len(game.Seats)
	game.Width = len(game.Seats[0])

	expecting := []int{37, 26}

	count1 := game.Run(false)
	if count1 != expecting[0] {
		t.Fatalf("Part 1 - Invalid count: %d", count1)
	}

	count2 := game.Run(true)
	if count2 != expecting[1] {
		t.Fatalf("Part 2 - Invalid count: %d", count2)
	}
}

func TestSolutions12(t *testing.T) {
	lines := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	expecting := []int{25, 286}

	ferry := &Ferry{
		Direction: 0,
		Position:  &Point{0, 0},
		Waypoint:  &Point{10, 1},
	}
	actions := []*FerryAction{}
	for _, line := range lines {
		action := ParseCourse(ferry, line)
		actions = append(actions, action)
	}

	distance1 := ferry.Distance()
	if distance1 != expecting[0] {
		t.Fatalf("Part 1 - Invalid count: %d", distance1)
	}

	ferry.Position.X = 0
	ferry.Position.Y = 0
	for _, action := range actions {
		ferry.Process(action)
	}

	distance2 := ferry.Distance()
	if distance2 != expecting[1] {
		t.Fatalf("Part 2 - Invalid count: %d", distance2)
	}
}

func TestSolution13(t *testing.T) {
	lines := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	expecting1 := []int{295}

	earliest := EarliestBus(lines[0], lines[1])
	if earliest != expecting1[0] {
		t.Fatalf("Part 1 - Invalid value: %d", earliest)
	}

	searchs := []string {
		"17,x,13,19",
		"67,7,59,61",
		"67,x,7,59,61",
		"67,7,x,59,61",
		"1789,37,47,1889",
		lines[1],
	}
	expecting2 := []int{3417, 754018, 779210, 1261476, 1202161486, 1068781}

	for s, search := range searchs {
		contest := Contest(search)
		if contest != expecting2[s] {
			t.Fatalf("Part 2 - Invalid value for %d: %d", s, contest)
		}
	}
}

func TestSolution14(t *testing.T) {
	expecting := []int{165, 208}
	lines := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	docker := NewDocker()
	for _, line := range lines {
		docker.Read1(line)
	}

	sum := docker.Sum()
	if sum != expecting[0] {
		t.Fatalf("Part 1 - Invalid value for %d", sum)
	}

	lines = []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	docker = NewDocker()
	for _, line := range lines {
		docker.Read2(line)
	}

	sum = docker.Sum()
	if sum != expecting[1] {
		t.Fatalf("Part 2 - Invalid value for %d", sum)
	}
}

func TestSolution15(t *testing.T) {
	/*
		Part 2 is disabled because it takes a long time to process
	*/

	expecting1 := []int{436,1,10,27,78,438,1836}
	//expecting2 := []int{175594,2578,3544142,261214,6895259,18,362}
	lines := []string{
		"0,3,6",
		"1,3,2",
		"2,1,3",
		"1,2,3",
		"2,3,1",
		"3,2,1",
		"3,1,2",
	}

	for l, line := range lines {
		game := ParseMemoryGame(line)
		part1 := game.Run(2020)
		if part1 != expecting1[l] {
			t.Fatalf("Part 1 - Invalid value for %d: %d", l, part1)
		}

		//part2 := game.Run(30000000)
		//if part2 != expecting2[l] {
		//	t.Fatalf("Part 2 - Invalid value for %d: %d", l, part2)
		//}
	}
}

func TestSolution16(t *testing.T) {
	expecting := []int{71}
	lines := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}

	ts := NewTicketSystem()
	for _, line := range lines {
		ts.Parse(line)
	}

	check := ts.Check()
	if check != expecting[0] {
		t.Fatalf("Part 1 - Invalid value:  %d", check)
	}
}
