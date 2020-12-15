package solutions

import (
	"github.com/aewens/aoc20/pkg/shared"
)

func init() {
	Map[15] = Solution15
}

type MemoryGame struct {
	Turn int
	Last int
	Seen map[int][]int
}

func NewMemoryGame() *MemoryGame {
	return &MemoryGame{
		Turn: 0,
		Seen: make(map[int][]int),
	}
}

func (game *MemoryGame) Next() {
	game.Turn = game.Turn + 1
}

func (game *MemoryGame) Add(value int) {
	game.Last = value
	turns, ok := game.Seen[game.Last]
	if !ok {
		turns = []int{}
	}
	game.Seen[game.Last] = append(turns, game.Turn)
}

func (game *MemoryGame) Step() {
	game.Next()
	turns, ok := game.Seen[game.Last]
	if !ok {
		panic("Something went very wrong")
	}

	if len(turns) == 1 {
		game.Add(0)
	} else {
		last := turns[len(turns)-1]
		prev := turns[len(turns)-2]
		game.Add(last - prev)
	}
}

func (game *MemoryGame) Run(stop int) int {
	for {
		if game.Turn == stop {
			break
		}
		game.Step()
	}
	return game.Last
}

func ParseMemoryGame(line string) *MemoryGame {
	game := NewMemoryGame()
	for _, value := range Separate(line, ",") {
		number := shared.StringToInt(value)
		game.Next()
		game.Seen[number] = []int{game.Turn}
		game.Last = number
	}

	return game
}

func Solution15(lines chan string) {
	game := ParseMemoryGame(<-lines)
	part1 := game.Run(2020)
	Display(1, part1)
}
