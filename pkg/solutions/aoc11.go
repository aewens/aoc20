package solutions

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

func (game *GameOfSeats) Step() bool {
	game.Used = 0
	changed := false
	next := make(map[int]map[int]int)
	for y := 0; y < game.Height; y++ {
		_, ok := next[y]
		if !ok {
			next[y] = make(map[int]int)
		}
		for x := 0; x < game.Width; x++ {
			seat := game.Seats[y][x]
			adjacents := game.Adjacent(y, x)
			if seat == empty && adjacents == 0 {
				next[y][x] = occupied
				changed = true
			} else if seat == occupied && adjacents >= 4 {
				next[y][x] = empty
				changed = true
			} else {
				next[y][x] = seat
			}

			if next[y][x] == occupied {
				game.Used = game.Used + 1
			}
		}
	}

	if changed {
		game.Seats = next
	}

	return changed
}

func (game *GameOfSeats) Run() int {
	for {
		if !game.Step() {
			break
		}
	}

	return game.Used
}

func Solution11(lines chan string) {
	game := &GameOfSeats{
		Seats: make(map[int]map[int]int),
	}
	for line := range lines {
		ParseGame(game, line)
	}
	game.Height = len(game.Seats)
	game.Width = len(game.Seats[0])

	Display(1, game.Run())
}
