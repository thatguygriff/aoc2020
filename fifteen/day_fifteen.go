package fifteen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	starting  []int
	numbers   []int
	spoken    map[int]int
	lastIndex map[int]int
}

func (g *game) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")

		for _, i := range input {
			v, err := strconv.Atoi(i)
			if err != nil {
				return err
			}

			g.starting = append(g.starting, v)
		}
	}

	return nil
}

func (g *game) valueAt(index int) int {
	g.spoken = make(map[int]int, index)
	g.lastIndex = make(map[int]int, index)
	for i, n := range g.starting {
		if i != len(g.starting)-1 {
			g.spoken[n]++
		}
		g.lastIndex[n] = i
		g.numbers = append(g.numbers, n)
	}

	for i := len(g.starting); i < index; i++ {
		last := g.numbers[i-1]
		g.spoken[last]++
		lastIndex := g.lastIndex[last]
		g.lastIndex[last] = i - 1

		next := 0

		if g.spoken[last] > 1 {
			next = i - 1 - lastIndex
		}

		g.numbers = append(g.numbers, next)
	}

	if len(g.numbers) != index {
		return -1
	}

	return g.numbers[index-1]
}

// PartOne What is the 2020th number given my input
func PartOne() string {
	g := game{}
	if err := g.load("fifteen/input.txt"); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The 2020th number is %d", g.valueAt(2020))
}
