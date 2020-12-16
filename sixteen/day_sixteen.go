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
	fields []string
}

type constraint struct {
	min, max int
}

type ticketScanner struct {
	rules  map[string][]constraint
	me     *ticket
	nearby []ticket
}

func (ts *ticketScanner) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	ts.rules = make(map[string][]constraint)
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

func (ts *ticketScanner) nearbyErrorRate() int {
	errorRate := 0
	for _, t := range ts.nearby {
		for _, value := range t.values {
			valid := false
			for _, rule := range ts.rules {
				for _, r := range rule {
					if value >= r.min && value <= r.max {
						valid = true
						break
					}
				}
				if valid {
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

// PartOne What is my nearby ticket scanning error rate
func PartOne() string {
	ts := ticketScanner{}
	if err := ts.load("sixteen/input.txt"); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The nearby scanning error rate is %d", ts.nearbyErrorRate())
}
