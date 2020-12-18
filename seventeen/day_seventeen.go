package seventeen

import (
	"bufio"
	"fmt"
	"os"
)

type pocket struct {
	dimension [][][]bool
}

func (p *pocket) initialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	z := 0
	p.dimension = [][][]bool{}
	p.dimension = append(p.dimension, [][]bool{})
	for scanner.Scan() {
		p.dimension[z] = append(p.dimension[z], []bool{})
		for _, c := range scanner.Text() {
			switch string(c) {
			case "#":
				p.dimension[z][y] = append(p.dimension[z][y], true)
			case ".":
				p.dimension[z][y] = append(p.dimension[z][y], false)
			}
		}
		y++
	}

	return nil
}

func (p *pocket) active() int {
	count := 0

	for z := 0; z < len(p.dimension); z++ {
		for y := 0; y < len(p.dimension[z]); y++ {
			for x := 0; x < len(p.dimension[z][y]); x++ {
				if p.dimension[z][y][x] {
					count++
				}
			}
		}
	}

	return count
}

func activation(p pocket, x, y, z int) bool {
	activeNeighbours := 0
	selfActive := false
	for zIndex := z - 1; zIndex <= z+1; zIndex++ {
		if zIndex < 0 || zIndex >= len(p.dimension) {
			continue
		}
		for yIndex := y - 1; yIndex <= y+1; yIndex++ {
			if yIndex < 0 || yIndex >= len(p.dimension[zIndex]) {
				continue
			}
			for xIndex := x - 1; xIndex <= x+1; xIndex++ {
				if xIndex < 0 || xIndex >= len(p.dimension[zIndex][yIndex]) {
					continue
				}

				if xIndex == x && yIndex == y && zIndex == z {
					selfActive = p.dimension[zIndex][yIndex][xIndex]
					continue
				}

				if p.dimension[zIndex][yIndex][xIndex] {
					activeNeighbours++
				}
			}
		}
	}

	if selfActive {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			return true
		}
	} else {
		if activeNeighbours == 3 {
			return true
		}
	}

	return false
}

func (p *pocket) print() {
	for z := 0; z < len(p.dimension); z++ {
		fmt.Printf("\nz=%d\n", z)
		for y := 0; y < len(p.dimension[z]); y++ {
			yString := ""
			for x := 0; x < len(p.dimension[z][y]); x++ {
				if p.dimension[z][y][x] {
					yString += "#"
				} else {
					yString += "."
				}
			}
			fmt.Println(yString)
		}
	}
}

func simulate(start pocket, count int) pocket {
	simulated := start

	for round := 0; round < count; round++ {
		next := pocket{
			dimension: make([][][]bool, len(simulated.dimension)+2),
		}

		for z := -1; z < len(simulated.dimension)+1; z++ {
			next.dimension[z+1] = make([][]bool, len(simulated.dimension[0])+2)
			for y := -1; y < len(simulated.dimension[0])+1; y++ {
				next.dimension[z+1][y+1] = make([]bool, len(simulated.dimension[0][0])+2)
				for x := -1; x < len(simulated.dimension[0][0])+1; x++ {
					next.dimension[z+1][y+1][x+1] = activation(simulated, x, y, z)
				}
			}
		}

		simulated = next
	}

	return simulated
}

// PartOne How many active cubes are left after 6 cycles
func PartOne() string {
	p := pocket{}
	if err := p.initialize("seventeen/input.txt"); err != nil {
		return err.Error()
	}

	p = simulate(p, 6)
	return fmt.Sprintf("There are %d active cubes after 6 rounds", p.active())
}
