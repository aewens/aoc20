package main

import (
	"flag"
	"fmt"

	"github.com/aewens/aoc20/pkg/shared"
	"github.com/aewens/aoc20/pkg/solutions"
)

type FlagState struct {
	Problem  int
	Override string
}

func parseFlags() *FlagState {
	problemFlag := flag.Int("p", -1, "Problem index to run")
	overrideFlag := flag.String("o", "", "Override problem file with own input")
	flag.Parse()

	if *problemFlag == -1 {
		panic("Missing problem flag")
	}

	if *problemFlag < 0 && *problemFlag > 25 {
		panic("Problem index is not in-between 1 and 25")
	}

	state := &FlagState{
		Problem:  *problemFlag,
		Override: *overrideFlag,
	}

	return state
}

func main() {
	//defer shared.Cleanup()
	shared.HandleSigterm()

	state := parseFlags()
	problem := state.Problem
	lines := make(chan string)

	if len(state.Override) == 0 {
		inputFile := fmt.Sprintf("etc/aoc%d.txt", problem)

		go solutions.ReadLines(inputFile, lines)
	} else {
		go func() {
			lines <- state.Override
		}()
	}

	solutions.Map[problem](lines)
}
