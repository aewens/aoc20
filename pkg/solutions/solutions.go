package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Solutions map[int]func(chan string)

var Map Solutions = make(Solutions)

func ReadLines(path string, lines chan string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines <- line
	}

	close(lines)

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Display(answer int, text interface{}) {
	fmt.Printf("[%d] %#v\n", answer, text)
}

func Input(prompt string) string {
	var input string

	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	return input
}

func FromCSV(source string) []string {
	return strings.Split(source, ",")
}

func Separate(source string, delim string) []string {
	return strings.Split(source, delim)
}

func Parameters(source string) []string {
	return strings.Fields(source)
}
