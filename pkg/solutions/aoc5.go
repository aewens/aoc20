package solutions

//import (
//	"github.com/aewens/aoc20/pkg/shared"
//)

func init() {
	Map[5] = Solution5
}

type Seat struct {
	Row    int
	Column int
	Id     int
}

type SeatRange struct {
	Row    *MinMax
	Column *MinMax
}

type Partition struct {
	Row    []bool
	Column []bool
}

func ParsePartition(line string) *Partition {
	partition := &Partition{
		Row:    []bool{},
		Column: []bool{},
	}
	for _, char := range line {
		letter := string(char)
		switch letter {
		case "F":
			partition.Row = append(partition.Row, true)
		case "B":
			partition.Row = append(partition.Row, false)
		case "L":
			partition.Column = append(partition.Column, true)
		case "R":
			partition.Column = append(partition.Column, false)
		default:
			panic("Invalid instruction: " + letter)
		}
	}

	return partition
}

func SearchSeat(partition *Partition) *Seat {
	seat := &Seat{}
	seatRange := SeatRange{
		Row:    &MinMax{0, 127},
		Column: &MinMax{0, 7},
	}
	for r, rowSide := range partition.Row {
		rowHalf := (seatRange.Row.Max+seatRange.Row.Min)/2
		if rowSide {
			if r == len(partition.Row)-1 {
				seat.Row = seatRange.Row.Min
			} else {
				seatRange.Row.Max = rowHalf
			}
		} else {
			if r == len(partition.Row)-1 {
				seat.Row = seatRange.Row.Max
			} else {
				seatRange.Row.Min = rowHalf+1
			}
		}
	}
	for r, colSide := range partition.Column {
		colHalf := (seatRange.Column.Max+seatRange.Column.Min)/2
		if colSide {
			if r == len(partition.Column)-1 {
				seat.Column = seatRange.Column.Min
			} else {
				seatRange.Column.Max = colHalf
			}
		} else {
			if r == len(partition.Column)-1 {
				seat.Column = seatRange.Column.Max
			} else {
				seatRange.Column.Min = colHalf+1
			}
		}
	}

	seat.Id = seat.Row * 8 + seat.Column
	return seat
}

func Solution5(lines chan string) {
	partitions := []*Partition{}
	for line := range lines {
		partitions = append(partitions, ParsePartition(line))
	}

	highestId := 0
	rows := &MinMax{-1, -1}
	columns := &MinMax{-1, -1}
	known := &MinMax{-1, -1}
	sum := 0
	seats := make(map[int]map[int]*Seat)
	for _, partition := range partitions {
		seat := SearchSeat(partition)
		if seat.Id > highestId {
			highestId = seat.Id
		}

		sum = sum + seat.Id
		if known.Min == -1 || seat.Id < known.Min {
			known.Min = seat.Id
		}

		if known.Max == -1 || seat.Id > known.Max {
			known.Max = seat.Id
		}

		seatsRow, ok := seats[seat.Row]
		if !ok {
			seats[seat.Row] = make(map[int]*Seat)
			seatsRow = seats[seat.Row]
		}

		seatsRow[seat.Column] = seat

		if rows.Min == -1 || seat.Row < rows.Min {
			rows.Min = seat.Row
		}

		if rows.Max == -1 || seat.Row > rows.Max {
			rows.Max = seat.Row
		}

		if columns.Min == -1 || seat.Column < columns.Min {
			columns.Min = seat.Column
		}

		if columns.Max == -1 || seat.Column > columns.Max {
			columns.Max = seat.Column
		}
	}

	Display(1, highestId)

	first := false
	found := false
	for y := rows.Min; y <= rows.Max; y++ {
		for x := columns.Min; x <= columns.Max; x++ {
			_, ok := seats[y][x]
			if ok && !first {
				first = true
				continue
			}
			if !ok && first {
				Display(2, y * 8 + x)
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// fancy arithmetic version of part 2
	count := known.Max-known.Min+1
	Display(3, count*(known.Min+known.Max)/2-sum)
}
