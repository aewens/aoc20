package solutions

import (
	"regexp"

	"github.com/aewens/aoc20/pkg/shared"
)

type Passport struct {
	Valid  bool
	Fields map[string]string
}

func init() {
	Map[4] = Solution4
}

func validYear(value string, minYear int, maxYear int) bool {
	if len(value) != 4 {
		return false
	}
	intValue := shared.StringToInt(value)
	return intValue >= minYear && intValue <= maxYear
}

func CheckPassportField(key string, value string) bool {
	switch key {
	case "byr":
		return validYear(value, 1920, 2002)
	case "iyr":
		return validYear(value, 2010, 2020)
	case "eyr":
		return validYear(value, 2020, 2030)
	case "hgt":
		if len(value) < 3 {
			return false
		}
		suffix := value[len(value)-2:]
		prefix := shared.StringToInt(value[:len(value)-2])
		switch suffix {
		case "cm":
			return prefix >= 150 && prefix <= 193
		case "in":
			return prefix >= 59 && prefix <= 76
		default:
			return false
		}
	case "hcl":
		match, err := regexp.MatchString("#[0-9a-f]{6}", value)
		if err != nil {
			panic(err)
		}
		return match
	case "ecl":
		switch value {
		case "amb":
			fallthrough
		case "blu":
			fallthrough
		case "brn":
			fallthrough
		case "gry":
			fallthrough
		case "grn":
			fallthrough
		case "hzl":
			fallthrough
		case "oth":
			return true
		default:
			return false
		}
	case "pid":
		match, err := regexp.MatchString("[0-9]{9}", value)
		if err != nil {
			panic(err)
		}
		return match
	case "cid":
		return true
	default:
		return false
	}
}

func ParsePassports(lines []string) []*Passport {
	passports := []*Passport{}
	passport := &Passport{
		Valid: false,
		Fields: make(map[string]string),
	}

	for _, line := range lines {
		if len(line) == 0 {
			if len(passport.Fields) == 8 {
				passport.Valid = true
			} else if len(passport.Fields) == 7 {
				_, ok := passport.Fields["cid"]
				passport.Valid = !ok
			} else {
				passport.Valid = false
			}

			passports = append(passports, passport)
			passport = &Passport{
				Fields: make(map[string]string),
			}
			continue
		}

		fields := Separate(line, " ")
		for _, field := range fields {
			pair := Separate(field, ":")
			passport.Fields[pair[0]] = pair[1]
		}
	}
	return passports
}

func Solution4(lines chan string) {
	// This bit is done to make testing eaiser
	data := []string{}
	for line := range lines {
		data = append(data, line)
	}

	passports := ParsePassports(data)

	valid1 := 1
	valid2 := 0
	for _, passport := range passports {
		if !passport.Valid {
			continue
		}

		valid1 = valid1 + 1
		invalid := false
		for key, value := range passport.Fields {
			if !CheckPassportField(key, value) {
				invalid = true
				break
			}
		}

		if !invalid {
			valid2 = valid2 + 1
		}
	}

	Display(1, valid1)
	Display(2, valid2)
}
