package seventeen

import (
	"bufio"
	"fmt"
	"os"
)

type pocket struct {
	dimension [][][][]bool
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
	w := 0
	p.dimension = [][][][]bool{}
	p.dimension = append(p.dimension, [][][]bool{})
	p.dimension[w] = append(p.dimension[w], [][]bool{})
	for scanner.Scan() {
		p.dimension[w][z] = append(p.dimension[w][z], []bool{})
		for _, c := range scanner.Text() {
			switch string(c) {
			case "#":
				p.dimension[w][z][y] = append(p.dimension[w][z][y], true)
			case ".":
				p.dimension[w][z][y] = append(p.dimension[w][z][y], false)
			}
		}
		y++
	}

	return nil
}

func (p *pocket) active() int {
	count := 0

	for w := 0; w < len(p.dimension); w++ {
		for z := 0; z < len(p.dimension[w]); z++ {
			for y := 0; y < len(p.dimension[w][z]); y++ {
				for x := 0; x < len(p.dimension[w][z][y]); x++ {
					if p.dimension[w][z][y][x] {
						count++
					}
				}
			}
		}
	}

	return count
}

func activation(p pocket, x, y, z, w int) bool {
	activeNeighbours := 0
	selfActive := false
	for wIndex := w - 1; wIndex <= w+1; wIndex++ {
		if wIndex < 0 || wIndex >= len(p.dimension) {
			continue
		}
		for zIndex := z - 1; zIndex <= z+1; zIndex++ {
			if zIndex < 0 || zIndex >= len(p.dimension[wIndex]) {
				continue
			}
			for yIndex := y - 1; yIndex <= y+1; yIndex++ {
				if yIndex < 0 || yIndex >= len(p.dimension[wIndex][zIndex]) {
					continue
				}
				for xIndex := x - 1; xIndex <= x+1; xIndex++ {
					if xIndex < 0 || xIndex >= len(p.dimension[wIndex][zIndex][yIndex]) {
						continue
					}

					if xIndex == x && yIndex == y && zIndex == z && wIndex == w {
						selfActive = p.dimension[wIndex][zIndex][yIndex][xIndex]
						continue
					}

					if p.dimension[wIndex][zIndex][yIndex][xIndex] {
						activeNeighbours++
					}
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
	for w := 0; w < len(p.dimension); w++ {
		for z := 0; z < len(p.dimension[w]); z++ {
			fmt.Printf("\nz=%d, w=%d\n", z, w)
			for y := 0; y < len(p.dimension[w][z]); y++ {
				yString := ""
				for x := 0; x < len(p.dimension[w][z][y]); x++ {
					if p.dimension[w][z][y][x] {
						yString += "#"
					} else {
						yString += "."
					}
				}
				fmt.Println(yString)
			}
		}
	}
}

func simulate(start pocket, count int) pocket {
	simulated := start

	for round := 0; round < count; round++ {
		next := pocket{
			dimension: make([][][][]bool, len(simulated.dimension)+2),
		}

		for w := -1; w < len(simulated.dimension)+1; w++ {
			next.dimension[w+1] = make([][][]bool, len(simulated.dimension[0])+2)
			for z := -1; z < len(simulated.dimension[0])+1; z++ {
				next.dimension[w+1][z+1] = make([][]bool, len(simulated.dimension[0][0])+2)
				for y := -1; y < len(simulated.dimension[0][0])+1; y++ {
					next.dimension[w+1][z+1][y+1] = make([]bool, len(simulated.dimension[0][0][0])+2)
					for x := -1; x < len(simulated.dimension[0][0][0])+1; x++ {
						next.dimension[w+1][z+1][y+1][x+1] = activation(simulated, x, y, z, w)
					}
				}
			}
		}

		simulated = next
	}

	return simulated
}

// PartTwo How many active cubes are left after 6 cycles
func PartTwo() string {
	p := pocket{}
	if err := p.initialize("seventeen/input.txt"); err != nil {
		return err.Error()
	}

	p = simulate(p, 6)
	return fmt.Sprintf("There are %d active cubes after 6 rounds", p.active())
}
