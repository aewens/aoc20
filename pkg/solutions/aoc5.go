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
	for _, partition := range partitions {
		seat := SearchSeat(partition)
		if seat.Id > highestId {
			highestId = seat.Id
		}
	}

	Display(1, highestId)
}
