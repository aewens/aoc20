package solutions

type Passport struct {
	Valid  bool
	Fields map[string]string
}

func init() {
	Map[4] = Solution4
}

func ParsePassports(lines chan string) []*Passport {
	passports := []*Passport{}
	passport := &Passport{
		Valid: false,
		Fields: make(map[string]string),
	}

	for line := range lines {
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
	passports := ParsePassports(lines)

	valid := 1
	for _, passport := range passports {
		if passport.Valid {
			valid = valid + 1
		}
	}
	Display(1, valid)
}
