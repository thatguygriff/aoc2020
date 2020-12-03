package three

import (
	"bufio"
	"fmt"
	"os"
)

type forest struct {
	trees       [][]bool
	rows, width int
}

func (f *forest) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if f.rows == 0 {
			f.width = len(scanner.Text())
		}

		row := make([]bool, f.width)
		for position, possibleTree := range scanner.Text() {
			row[position] = (string(possibleTree) == "#")
		}
		f.trees = append(f.trees, row)
		f.rows++
	}

	return nil
}

func (f *forest) toboggan(right, down int) int {
	x, y, trees := 0, 0, 0
	for y < len(f.trees)-1 {
		y = y + down
		x = x + right
		if x > (f.width - 1) {
			x = x % f.width
		}

		if f.trees[y][x] {
			trees++
		}
	}

	return trees
}

func PartOne() string {
	forest := forest{}
	forest.load("three/trees.txt")
	trees := forest.toboggan(3, 1)
	return fmt.Sprintf("Found %d trees on right 3, down 1", trees)
}

func PartTwo() string {
	forest := forest{}
	forest.load("three/trees.txt")
	product := forest.toboggan(1, 1) *
		forest.toboggan(3, 1) *
		forest.toboggan(5, 1) *
		forest.toboggan(7, 1) *
		forest.toboggan(1, 2)
	return fmt.Sprintf("Found product %d of trees over 5 runs", product)
}
