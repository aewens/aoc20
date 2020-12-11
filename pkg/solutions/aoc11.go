package solutions

import (
	"time"
)

func init() {
	Map[11] = Solution11
}

const (
	floor int = iota
	empty
	occupied
)

type GameOfSeats struct {
	Width  int
	Height int
	Seats  map[int]map[int]int
	Copy   map[int]map[int]int
	Used   int
}

func ParseGame(game *GameOfSeats, line string) {
	row := make(map[int]int)
	for l, letter := range line {
		switch letter {
		case '.':
			row[l] = floor
		case 'L':
			row[l] = empty
		default:
			panic("Invalid seat")
		}
	}
	game.Seats[len(game.Seats)] = row
	game.Copy[len(game.Copy)] = row
}

func (game *GameOfSeats) Adjacent(y int, x int) int {
	adjacents := 0
	for dy := -1; dy <= 1; dy++ {
		row, ok := game.Seats[y+dy]
		if !ok {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			seat, ok := row[x+dx]
			if !ok {
				continue
			}

			if seat == occupied {
				adjacents = adjacents + 1
			}
		}
	}

	return adjacents
}

func (game *GameOfSeats) Simple(y int, x int) int {
	seat := game.Seats[y][x]
	adjacents := game.Adjacent(y, x)
	if seat == empty && adjacents == 0 {
		return occupied
	}
	if seat == occupied && adjacents >= 4 {
		return empty
	}
	return seat
}

func (game *GameOfSeats) Visibles(y int, x int) int {
	visibles := 0
	for oy := -1; oy <= 1; oy++ {
		for ox := -1; ox <= 1; ox++ {
			if oy == 0 && ox == 0 {
				continue
			}

			dy := oy
			dx := ox
			for {
				seat, ok := game.Seats[y+dy][x+dx]
				if !ok {
					break
				}

				if seat == floor {
					dy = dy + oy
					dx = dx + ox
					continue
				}

				if seat == occupied {
					visibles = visibles + 1
				}
				break
			}
		}
	}
	return visibles
}

func (game *GameOfSeats) Complex(y int, x int) int {
	seat := game.Seats[y][x]
	visibles := game.Visibles(y, x)
	if seat == empty && visibles == 0 {
		return occupied
	}
	if seat == occupied && visibles >= 5 {
		return empty
	}
	return seat
}

func (game *GameOfSeats) Step(part2 bool) bool {
	game.Used = 0
	changed := false
	next := make(map[int]map[int]int)
	for y := 0; y < game.Height; y++ {
		_, ok := next[y]
		if !ok {
			next[y] = make(map[int]int)
		}
		for x := 0; x < game.Width; x++ {
			var seat int
			if !part2 {
				seat = game.Simple(y, x)
			} else {
				seat = game.Complex(y, x)
			}
			next[y][x] = seat

			if seat != game.Seats[y][x] {
				changed = true
			}

			if seat == occupied {
				game.Used = game.Used + 1
			}
		}
	}

	if changed {
		game.Seats = next
	}

	//if part2 {
	//	game.Display()
	//}

	return changed
}

func (game *GameOfSeats) Run(part2 bool) int {
	game.Seats = game.Copy
	for {
		if !game.Step(part2) {
			break
		}
	}

	return game.Used
}

func (game *GameOfSeats) Display() {
	Clear()
	for y := 0; y < game.Height; y++ {
		row := ""
		for x := 0; x < game.Width; x++ {
			switch game.Seats[y][x] {
			case floor:
				row = row + "."
			case empty:
				row = row + "L"
			case occupied:
				row = row + "#"
			}
		}
		Display(0, row)
	}
	time.Sleep(2 * time.Second)
}

func Solution11(lines chan string) {
	game := &GameOfSeats{
		Seats: make(map[int]map[int]int),
		Copy: make(map[int]map[int]int),
	}
	for line := range lines {
		ParseGame(game, line)
	}
	game.Height = len(game.Seats)
	game.Width = len(game.Seats[0])

	Display(1, game.Run(false))
	Display(2, game.Run(true))
}
