package shared

import (
	"strconv"
)

func StringToInt(convert string) int {
	value, err := strconv.Atoi(convert)
	if err != nil {
		panic(err)
	}

	return value
}

func RuneToInt(convert rune) int {
	return int(convert - '0')
}

func IntToString(convert int) string {
	value := strconv.Itoa(convert)
	return value
}
