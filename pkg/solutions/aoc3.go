package solutions

//import (
//	"github.com/aewens/aoc20/pkg/shared"
//)

func init() {
	Map[3] = Solution3
}

type Point struct {
	X int
	Y int
}

type TreeMap struct {
	Width    int
	Height   int
	Pattern  [][]bool
	Toboggan *Point
	Vector   *Point
	TreesHit int
}

func (tm *TreeMap) Init() {
	tm.Width = len(tm.Pattern[0])
	tm.Height = len(tm.Pattern)
}

func (tm *TreeMap) Reset() {
	tm.TreesHit = 0
	tm.Toboggan.X = 0
	tm.Toboggan.Y = 0
}

func (tm *TreeMap) Upgrade(vector *Point) {
	tm.Vector = vector
	tm.Reset()
}

func (tm *TreeMap) Step() bool {
	y := tm.Toboggan.Y
	if y >= tm.Height {
		return true
	}

	x := tm.Toboggan.X % tm.Width
	if tm.Pattern[y][x] {
		tm.TreesHit = tm.TreesHit + 1
	}

	tm.Toboggan.X = tm.Toboggan.X + tm.Vector.X
	tm.Toboggan.Y = tm.Toboggan.Y + tm.Vector.Y
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
		Toboggan: &Point{0, 0},
		Vector: &Point{3, 1},
	}

	for line := range lines {
		ParsePattern(treeMap, line)
	}

	treeMap.Init()
	treeMap.Descend()
	Display(1, treeMap.TreesHit)

	hits := treeMap.TreesHit
	vectors := []*Point{
		{1, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, vector := range vectors {
		treeMap.Upgrade(vector)
		treeMap.Descend()
		hits = hits * treeMap.TreesHit
	}

	Display(2, hits)
}
