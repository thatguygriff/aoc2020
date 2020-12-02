package two

import (
	"bufio"
	"fmt"
	"os"
)

type passwordPolicy struct {
	min      int
	max      int
	mustHave string
}

type passwordEntry struct {
	policy   passwordPolicy
	password string
}

type db struct {
	passwords []passwordEntry
}

func (d *db) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var min, max int
		var mustHave, p string

		count, err := fmt.Sscanf(scanner.Text(), "%d-%d %1s: %s", &min, &max, &mustHave, &p)
		if err != nil {
			return err
		} else if count != 4 {
			return fmt.Errorf("Unable to parse entry %q", scanner.Text())
		}

		e := passwordEntry{
			policy: passwordPolicy{
				min:      min,
				max:      max,
				mustHave: mustHave,
			},
			password: p,
		}
		d.passwords = append(d.passwords, e)
	}

	return nil
}

func validate(password string, policy passwordPolicy) bool {
	instances := 0
	for _, letter := range password {
		if string(letter) == policy.mustHave {
			instances++
		}
	}

	return (instances >= policy.min && instances <= policy.max)
}

func (d *db) validate() int {
	valid := 0
	for _, entry := range d.passwords {
		if validate(entry.password, entry.policy) {
			valid++
		}
	}

	return valid
}

// PartOne Find how many passwords are valid according to their policy
func PartOne() string {
	database := db{}
	database.load("two/passwords.txt")
	valid := database.validate()
	return fmt.Sprintf("Found %d valid passwords", valid)
}

// PartTwo
func PartTwo() string {
	return ""
}
