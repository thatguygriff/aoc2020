package twentythree

import (
	"fmt"
	"strconv"
)

type cup struct {
	id   int
	next *cup
}

type game struct {
	currentCup *cup
	cups       int
	index      map[int]*cup
}

func (g *game) load(input string) error {
	if len(input) < 1 {
		return fmt.Errorf("Cannot parse empty input")
	}
	g.cups = len(input)

	g.index = map[int]*cup{}
	g.currentCup = &cup{}
	cp := g.currentCup
	for i, c := range input {
		cupNumber, err := strconv.Atoi(string(c))
		if err != nil {
			return err
		}

		cp.id = cupNumber
		g.index[cp.id] = cp

		if i != g.cups-1 {
			cp.next = &cup{}
			cp = cp.next
		} else {
			cp.next = g.currentCup
		}
	}

	return nil
}

func (g *game) milliload(input string) error {
	if err := g.load(input); err != nil {
		return err
	}

	start := g.currentCup.next
	for start.next != g.currentCup {
		start = start.next
	}

	for g.cups != 1000000 {
		start.next = &cup{
			id: g.cups + 1,
		}
		g.index[start.next.id] = start.next
		g.cups++
		start = start.next
	}
	start.next = g.currentCup

	return nil
}

func destination(cups, possible, missing1, missing2, missing3 int) int {
	d := possible
	if d <= 0 {
		d = cups
	}

	if d == missing1 || d == missing2 || d == missing3 {
		return destination(cups, d-1, missing1, missing2, missing3)
	}

	return d
}

func (g *game) play(moves int, output bool) string {
	for i := 0; i < moves; i++ {
		// pickup cups
		pickup := g.currentCup.next
		g.currentCup.next = g.currentCup.next.next.next.next

		// find destination cup
		dp := g.index[destination(g.cups, g.currentCup.id-1, pickup.id, pickup.next.id, pickup.next.next.id)]

		// place cups
		pickup.next.next.next = dp.next
		dp.next = pickup

		g.currentCup = g.currentCup.next
	}

	if output {
		return g.output()
	}

	return ""
}

func (g *game) output() string {
	cp := g.index[1].next
	output := ""
	for g.index[1] != cp {
		output += fmt.Sprintf("%d", cp.id)
		cp = cp.next
	}

	return output
}

func (g *game) starProduct() int {
	return g.index[1].next.id * g.index[1].next.next.id
}

// PartOne What are the labels on the cups after cup 1
func PartOne() string {
	g := game{}
	g.load("418976235")
	return g.play(100, true)
}

// PartTwo What is the product of the two cups clockwise of cup 1
func PartTwo() string {
	g := game{}
	g.milliload("418976235")
	g.play(10000000, false)
	return fmt.Sprintf("The start product is %d", g.starProduct())
}
