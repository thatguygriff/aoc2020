package twentythree

import (
	"fmt"
	"strconv"
)

type game struct {
	cups       []int
	currentCup int
	highestCup int
}

func (g *game) load(input string) error {
	for _, c := range input {
		cupNumber, err := strconv.Atoi(string(c))
		if err != nil {
			return err
		}
		if cupNumber > g.highestCup {
			g.highestCup = cupNumber
		}
		g.cups = append(g.cups, cupNumber)
	}

	if len(g.cups) > 0 {
		g.currentCup = g.cups[0]
	}

	return nil
}

func (g *game) play(moves int) string {
	for i := 0; i < moves; i++ {
		// pickup cups
		pickup := g.cups[1:4]
		remaining := []int{g.cups[0]}
		remaining = append(remaining, g.cups[4:]...)
		g.cups = remaining

		// find destination cup
		destination := false
		destinationIndex := -1
		t := g.currentCup - 1
		for !destination {
			if t <= 0 {
				t = g.highestCup
			}

			for i := 1; i < len(g.cups); i++ {
				if g.cups[i] == t {
					destinationIndex = i
					break
				}
			}

			if destinationIndex != -1 {
				destination = true
			}
			t--
		}

		// place cups
		newOrder := []int{}
		for _, c := range remaining[:destinationIndex+1] {
			newOrder = append(newOrder, c)
		}
		newOrder = append(newOrder, pickup...)
		for _, c := range remaining[destinationIndex+1:] {
			newOrder = append(newOrder, c)
		}
		g.cups = newOrder

		// Update current cup
		g.cups = g.cups[1:]
		g.cups = append(g.cups, g.currentCup)
		g.currentCup = g.cups[0]
	}

	return g.output()
}

func (g *game) output() string {
	state := []int{}
	for i, c := range g.cups {
		if c == 1 {
			state = append(state, g.cups[i+1:]...)
			state = append(state, g.cups[:i]...)
		}
	}
	output := ""
	for _, n := range state {
		output += fmt.Sprintf("%d", n)
	}

	return output
}

// PartOne What are the labels on the cups after cup 1
func PartOne() string {
	g := game{}
	g.load("418976235")
	return g.play(100)
}
