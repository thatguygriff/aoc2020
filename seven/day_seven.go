package seven

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bag struct {
	colour   string
	capacity map[string]int
}

type rules struct {
	bags []bag
}

func (r *rules) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var colour, c1, c2 string
		rule := strings.Split(scanner.Text(), " contain ")

		if len(rule) != 2 {
			return fmt.Errorf("Unable to parse %s", scanner.Text())
		}

		found, err := fmt.Sscanf(rule[0], "%s %s bags", &c1, &c2)
		if err != nil || found != 2 {
			return fmt.Errorf("Unable to parse colour from %s: %w", rule[0], err)
		}
		colour = fmt.Sprintf("%s %s", c1, c2)

		bag := bag{
			colour:   colour,
			capacity: make(map[string]int),
		}

		if rule[1] != "no other bags." {
			// Trim the trailing period
			raw := rule[1][:len(rule[1])-1]
			constraints := strings.Split(raw, ", ")

			for _, c := range constraints {
				var count int
				found, err := fmt.Sscanf(c, "%d %s %s bag", &count, &c1, &c2)
				if err != nil || found != 3 {
					return fmt.Errorf("Unable to parse colour from %s: %w", c, err)
				}
				colour = fmt.Sprintf("%s %s", c1, c2)

				bag.capacity[colour] = count
			}
		}

		r.bags = append(r.bags, bag)
	}

	return nil
}

func (r *rules) canContain(colours ...string) (int, []string) {
	holding := []string{}

	for _, c := range colours {
		for _, b := range r.bags {
			if _, canHold := b.capacity[c]; canHold {
				holding = append(holding, b.colour)
			}
		}
	}

	if len(holding) > 0 {
		_, additional := r.canContain(holding...)

		for _, c := range additional {
			newColour := true
			for _, e := range holding {
				if c == e {
					newColour = false
				}
			}
			if newColour {
				holding = append(holding, c)
			}
		}
	}

	return len(holding), holding
}

func (r *rules) get(colour string) *bag {
	for _, b := range r.bags {
		if b.colour == colour {
			return &b
		}
	}

	return nil
}

func (r *rules) canHold(colour string) int {
	count := 0

	b := r.get(colour)
	if b == nil {
		return count
	}

	for colour, capcity := range b.capacity {
		switch r.canHold(colour) {
		case 0:
			count += capcity
		default:
			count += r.canHold(colour)*capcity + capcity
		}
	}

	return count
}

// PartOne How many bags can hold at least 1 shiny gold
func PartOne() string {
	r := rules{}
	r.load("seven/rules.txt")

	count, _ := r.canContain("shiny gold")

	return fmt.Sprintf("There are %d bags that can hold shiny gold", count)
}

// PartTwo How many bags must a shiny gold bag hold
func PartTwo() string {
	r := rules{}
	r.load("seven/rules.txt")

	count := r.canHold("shiny gold")

	return fmt.Sprintf("There are %d bags at maximum inside a shiny gold bag", count)
}
