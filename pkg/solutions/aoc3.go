package solutions

//import (
//	"github.com/aewens/aoc20/pkg/shared"
//)

func init() {
	Map[3] = Solution3
}

type Toboggan struct {
	X int
	Y int
}

type TreeMap struct {
	Width    int
	Height   int
	Pattern  [][]bool
	Toboggan *Toboggan
	TreesHit int
}

func (tm *TreeMap) Init() {
	tm.Width = len(tm.Pattern[0])
	tm.Height = len(tm.Pattern)
}

func (tm *TreeMap) Step() bool {
	y := tm.Toboggan.Y
	if y == tm.Height {
		return true
	}

	x := tm.Toboggan.X % tm.Width
	if tm.Pattern[y][x] {
		tm.TreesHit = tm.TreesHit + 1
	}

	tm.Toboggan.X = tm.Toboggan.X + 3
	tm.Toboggan.Y = tm.Toboggan.Y + 1
	return false
}

func (tm *TreeMap) Descend() {
	for {
		if tm.Step() {
			break
		}
	}
}

func ParsePattern(treeMap *TreeMap, line string) {
	var tree rune = '#'

	pattern := []bool{}
	for _, block := range line {
		pattern = append(pattern, block == tree)
	}

	treeMap.Pattern = append(treeMap.Pattern, pattern)
}

func Solution3(lines chan string) {
	treeMap := &TreeMap{
		Pattern:  [][]bool{},
		Toboggan: &Toboggan{0, 0},
	}

	for line := range lines {
		ParsePattern(treeMap, line)
	}

	treeMap.Init()
	treeMap.Descend()

	Display(1, treeMap.TreesHit)
}
