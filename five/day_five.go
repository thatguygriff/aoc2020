package five

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pass struct {
	raw             string
	row, column, id int
}

type plane struct {
	passes []pass
}

func decode(pass string) (row, column, id int) {
	rowString := pass[:7]
	seatString := pass[7:]

	rowMin, rowMax := 0, 127
	for _, r := range rowString {
		diff := ((rowMax - rowMin) / 2) + 1
		switch string(r) {
		case "F":
			row = rowMin
			rowMax -= diff
		case "B":
			row = rowMax
			rowMin += diff
		}
	}

	colMin, colMax := 0, 7
	for _, r := range seatString {
		diff := ((colMax - colMin) / 2) + 1
		switch string(r) {
		case "L":
			column = colMin
			colMax -= diff
		case "R":
			column = colMax
			colMin += diff
		}
	}

	return row, column, (row*8 + column)
}

func (p *plane) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row, column, id := decode(scanner.Text())
		p.passes = append(p.passes, pass{
			raw:    scanner.Text(),
			row:    row,
			column: column,
			id:     id,
		})
	}

	return nil
}

// PartOne find the highest boarding pass id
func PartOne() string {
	plane := plane{}
	plane.load("five/passes.txt")
	highestID := 0
	for _, pass := range plane.passes {
		if pass.id > highestID {
			highestID = pass.id
		}
	}

	return fmt.Sprintf("The highest seat id is %d on existing passes", highestID)
}

// PartTwo Find my boarding pass id
func PartTwo() string {
	plane := plane{}
	plane.load("five/passes.txt")

	sort.Slice(plane.passes, func(i, j int) bool {
		return plane.passes[i].id < plane.passes[j].id
	})
	found := 0
	for i := 0; i < len(plane.passes)-1; i++ {
		if (plane.passes[i+1].id - plane.passes[i].id) == 2 {
			found = plane.passes[i].id + 1
			break
		}
	}

	return fmt.Sprintf("My passport id is %d", found)
}
