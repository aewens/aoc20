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
	Seen map[int]int
}

func NewMemoryGame() *MemoryGame {
	return &MemoryGame{
		Turn: 0,
		Last: -1,
		Seen: make(map[int]int),
	}
}

func (game *MemoryGame) Next() {
	game.Turn = game.Turn + 1
}

func (game *MemoryGame) Step() {
	/*
	T 0
	L -1
	S 
	P

	T 1
	L 0
	S -1:1
	P -1:0

	T 2
	L 3
	S -1:1 0:1
	P -1:0 0:1

	T 3
	L 6
	S -1:1 0:1 3:1
	P -1:0 0:1 3:2

	T 4
	L 0
	S -1:1 0:1 3:1 6:1
	P -1:0 0:1 3:2 6:3

	T 5
	L 3 (4-1)
	S -1:1 0:1 3:1 6:1
	P -1:0 0:4 3:2 6:3

	T 6
	L 3 (5-2)
	S -1:1 0:1 3:1 6:1
	P -1:0 0:4 3:5 6:3

	T 7
	L 1 (6-5)
	S -1:1 0:1 3:1 6:1
	P -1:0 0:4 3:6 6:3

	T 8
	L 0
	S -1:1 0:1 3:1 6:1 1:1
	P -1:0 0:4 3:6 6:3

	T 9
	L 4 (8-4)
	S -1:1 0:1 3:1 6:1 1:1
	P -1:0 0:4 3:6 6:3 1:8

	T 10
	L 0
	S -1:1 0:1 3:1 6:1 1:1 4:1
	P -1:0 0:4 3:6 6:3 1:8
	*/
	game.Next()
	//Display(-1, game.Turn)
	//Display(-2, game.Last)
	turn, ok := game.Seen[game.Last]
	if !ok {
		game.Seen[game.Last] = game.Turn-1
		game.Last = 0
		//Display(-3, game.Last)
		//Display(-4, game.Seen)
		//Display(-5, game.Prev)
		return
	}

	game.Seen[game.Last] = game.Turn-1
	game.Last = (game.Turn-1)-turn
	//Display(-3, game.Last)
	//Display(-4, game.Seen)
	//Display(-5, game.Prev)
}

func (game *MemoryGame) Run(stop int) int {
	for {
		if game.Turn == stop {
			break
		}
		game.Step()
		//Display(-9, "----")
	}
	return game.Last
}

func ParseMemoryGame(line string) *MemoryGame {
	game := NewMemoryGame()
	for _, value := range Separate(line, ",") {
		number := shared.StringToInt(value)
		game.Next()
		game.Seen[game.Last] = game.Turn-1
		game.Last = number
	}

	return game
}

func Solution15(lines chan string) {
	game := ParseMemoryGame(<-lines)
	part1 := game.Run(2020)
	Display(1, part1)

	part2 := game.Run(30000000)
	Display(2, part2)
}
