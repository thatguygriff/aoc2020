package six

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	questions []string
}

type flight struct {
	groups []group
}

func (g *group) uniqueQuestion() int {
	questions := map[string]bool{}

	for _, answers := range g.questions {
		for _, q := range answers {
			if !questions[string(q)] {
				questions[string(q)] = true
			}
		}
	}

	return len(questions)
}

func (g *group) consensusQuestions() int {
	questions := map[string][]bool{}

	for _, answers := range g.questions {
		for _, q := range answers {
			questions[string(q)] = append(questions[string(q)], true)
		}
	}

	consensus := 0
	for _, q := range questions {
		if len(q) == len(g.questions) {
			consensus++
		}
	}

	return consensus
}

func (f *flight) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var g *group
	for scanner.Scan() {
		if scanner.Text() == "" {
			// Blank line means we have moved on to the next passport
			if g != nil {
				f.groups = append(f.groups, *g)
				g = nil
			}
			continue
		}

		if g == nil {
			g = &group{}
		}

		g.questions = append(g.questions, scanner.Text())
	}
	// commit the last passport if there isn't a blank line
	if g != nil {
		f.groups = append(f.groups, *g)
	}

	return nil
}

// PartOne Find the sum of all the unique questions in all groups
func PartOne() string {
	f := flight{}
	f.load("six/customs.txt")
	sum := 0
	for _, g := range f.groups {
		sum += g.uniqueQuestion()
	}

	return fmt.Sprintf("The sum of all the unique questions in all groups is %d", sum)
}

// PartTwo Find the sum of all the consensus questions in all groups
func PartTwo() string {
	f := flight{}
	f.load("six/customs.txt")
	sum := 0
	for _, g := range f.groups {
		sum += g.consensusQuestions()
	}

	return fmt.Sprintf("The sum of all the consensus questions in all groups is %d", sum)
}
