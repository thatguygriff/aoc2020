package sixteen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	values []int
}

type constraint struct {
	min, max int
}

type constraints []constraint

type ticketScanner struct {
	rules  map[string]constraints
	me     *ticket
	nearby []ticket
	valid  []ticket

	fields []string
}

func (ts *ticketScanner) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	ts.rules = make(map[string]constraints)
	var rules, myticket bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if scanner.Text() == "your ticket:" {
			rules = true
			continue
		}

		if !rules {
			parts := strings.Split(scanner.Text(), ":")
			if len(parts) != 2 {
				return fmt.Errorf("Unable to parse rule %s", scanner.Text())
			}
			ranges := []constraint{}
			c := strings.Split(parts[1], " or ")
			for _, a := range c {
				b := constraint{}
				count, err := fmt.Sscanf(a, "%d-%d", &b.min, &b.max)
				if count != 2 || err != nil {
					return fmt.Errorf("Unable to parse %s", a)
				}
				ranges = append(ranges, b)
			}
			ts.rules[parts[0]] = ranges
			continue
		}

		if scanner.Text() == "nearby tickets:" {
			myticket = true
			continue
		}

		if !myticket {
			t := &ticket{}
			values := strings.Split(scanner.Text(), ",")
			for _, v := range values {
				intValue, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				t.values = append(t.values, intValue)
			}
			ts.me = t
			continue
		}

		t := ticket{}
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			intValue, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			t.values = append(t.values, intValue)
		}
		ts.nearby = append(ts.nearby, t)
	}
	return nil
}

func (c *constraints) valid(input int) bool {
	for _, r := range *c {
		if input >= r.min && input <= r.max {
			return true
		}
	}
	return false
}

func (ts *ticketScanner) nearbyErrorRate() int {
	errorRate := 0
	for _, t := range ts.nearby {
		for _, value := range t.values {
			valid := false
			for _, rule := range ts.rules {
				if rule.valid(value) {
					valid = true
					break
				}
			}

			if !valid {
				errorRate += value
			}
		}
	}
	return errorRate
}

func (ts *ticketScanner) scanValid() {
	for _, t := range ts.nearby {
		validTicket := true
		for _, v := range t.values {
			validValue := false
			for _, r := range ts.rules {
				if r.valid(v) {
					validValue = true
					break
				}
			}
			if !validValue {
				validTicket = false
				break
			}
		}
		if validTicket {
			ts.valid = append(ts.valid, t)
		}
	}
}

func (ts *ticketScanner) determineFields() {
	ts.fields = make([]string, len(ts.rules))

	possibilities := map[string][]int{}
	for name, r := range ts.rules {
		possibilities[name] = []int{}

		for i := 0; i < len(ts.fields); i++ {
			isValid := true
			for j := 0; j < len(ts.valid); j++ {
				if !r.valid(ts.valid[j].values[i]) {
					isValid = false
				}
			}
			if isValid {
				possibilities[name] = append(possibilities[name], i)
			}
		}
	}

	for i := 0; i < len(ts.fields); i++ {
		if ts.fields[i] != "" {
			continue
		}
		options := []string{}
		for name, indexes := range possibilities {
			for _, index := range indexes {
				if i == index {
					options = append(options, name)
				}
			}
		}

		if len(options) == 1 {
			ts.fields[i] = options[0]
			possibilities = removeIndex(i, options[0], possibilities)
			continue
		}

		min := len(ts.fields)
		minKey := ""
		for _, option := range options {
			if len(possibilities[option]) < min {
				min = len(possibilities[option])
				minKey = option
			}
		}
		ts.fields[i] = minKey
		possibilities = removeIndex(i, minKey, possibilities)
	}
}

func removeIndex(index int, used string, possibilities map[string][]int) map[string][]int {
	updatedPossibilites := make(map[string][]int)
	for key, indexes := range possibilities {
		if used == key {
			continue
		}
		updatedIndexes := []int{}
		for _, i := range indexes {
			if i != index {
				updatedIndexes = append(updatedIndexes, i)
			}
		}

		updatedPossibilites[key] = updatedIndexes
	}
	return updatedPossibilites
}

func (ts *ticketScanner) fieldProduct(prefix string) int {
	product := 1
	for i, f := range ts.fields {
		if strings.Contains(f, prefix) {
			product *= ts.me.values[i]
		}
	}

	return product
}

// PartOne What is my nearby ticket scanning error rate
func PartOne() string {
	ts := ticketScanner{}
	if err := ts.load("sixteen/input.txt"); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The nearby scanning error rate is %d", ts.nearbyErrorRate())
}

// PartTwo What is the product of all fields beginning with "departure"
func PartTwo() string {
	ts := ticketScanner{}
	if err := ts.load("sixteen/input.txt"); err != nil {
		return err.Error()
	}

	ts.scanValid()
	ts.determineFields()

	return fmt.Sprintf("The product of departure is %d", ts.fieldProduct("departure"))
}
