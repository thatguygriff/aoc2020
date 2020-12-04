package four

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (p *passport) validate(strict bool, optional ...string) bool {
	valid := true
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	filtered := []string{}
	// remove optional fields
	for _, f := range requiredFields {
		required := true
		for _, o := range optional {
			if f == o {
				required = false
				break
			}
		}
		if required {
			filtered = append(filtered, f)
		}
	}

	for _, f := range filtered {
		if strict {
			if !p.validateField(f) {
				valid = false
			}
		} else {
			if p.getField(f) == "" {
				valid = false
			}
		}
	}

	return valid
}

func (p *passport) validateField(code string) bool {
	switch code {
	case "byr":
		year, err := strconv.Atoi(p.byr)
		if err != nil {
			return false
		}
		return year >= 1920 && year <= 2002
	case "iyr":
		year, err := strconv.Atoi(p.iyr)
		if err != nil {
			return false
		}
		return year >= 2010 && year <= 2020
	case "eyr":
		year, err := strconv.Atoi(p.eyr)
		if err != nil {
			return false
		}
		return year >= 2020 && year <= 2030
	case "hgt":
		var height int
		var unit string
		count, err := fmt.Sscanf(p.hgt, "%d%s", &height, &unit)
		if count != 2 || err != nil {
			return false
		}
		if unit == "in" {
			return height >= 59 && height <= 76
		}
		return height >= 150 && height <= 193
	case "hcl":
		if len(p.hcl) != 7 {
			return false
		}
		var validColour = regexp.MustCompile(`^#[a-z0-9]+$`)
		return validColour.MatchString(p.hcl)
	case "ecl":
		switch p.ecl {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	case "pid":
		if len(p.pid) != 9 {
			return false
		}
		var validColour = regexp.MustCompile(`^[0-9]+$`)
		return validColour.MatchString(p.pid)
	default:
		return false
	}
}

func (p *passport) getField(code string) string {
	switch code {
	case "byr":
		return p.byr
	case "iyr":
		return p.iyr
	case "eyr":
		return p.eyr
	case "hgt":
		return p.hgt
	case "hcl":
		return p.hcl
	case "ecl":
		return p.ecl
	case "pid":
		return p.pid
	case "cid":
		return p.cid
	default:
		return ""
	}
}

func (p *passport) setField(code, value string) {
	switch code {
	case "byr":
		p.byr = value
	case "iyr":
		p.iyr = value
	case "eyr":
		p.eyr = value
	case "hgt":
		p.hgt = value
	case "hcl":
		p.hcl = value
	case "ecl":
		p.ecl = value
	case "pid":
		p.pid = value
	case "cid":
		p.cid = value
	default:
		return
	}
}

type customs struct {
	passports []passport
}

func (c *customs) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var p *passport
	for scanner.Scan() {
		if scanner.Text() == "" {
			// Blank line means we have moved on to the next passport
			if p != nil {
				c.passports = append(c.passports, *p)
				p = nil
			}
			continue
		}

		if p == nil {
			p = &passport{}
		}

		fields := strings.Split(scanner.Text(), " ")
		for _, f := range fields {
			var key, value string
			count, err := fmt.Sscanf(f, "%3s:%s", &key, &value)
			if count != 2 || err != nil {
				return fmt.Errorf("Unable to parse field %s: %w", f, err)
			}

			p.setField(key, value)
		}
	}
	// commit the last passport if there isn't a blank line
	if p != nil {
		c.passports = append(c.passports, *p)
	}

	return nil
}

func (c *customs) check(strict bool, optional ...string) int {
	valid := 0

	for _, p := range c.passports {
		if p.validate(strict, optional...) {
			valid++
		}
	}

	return valid
}

// PartOne find all the valid passports
func PartOne() string {
	customs := customs{}
	customs.load("four/passports.txt")
	validPassports := customs.check(false, "cid")

	return fmt.Sprintf("Found %d valid passports", validPassports)
}

// PartTwo Strictly find all the valid passports
func PartTwo() string {
	customs := customs{}
	customs.load("four/passports.txt")
	validPassports := customs.check(true, "cid")

	return fmt.Sprintf("Found %d strictly valid passports", validPassports)
}
