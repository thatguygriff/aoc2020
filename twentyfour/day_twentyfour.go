package twentyfour

import (
	"bufio"
	"fmt"
	"os"
)

type lobby struct {
	instructions []string
	tiles        map[string]bool
}

func (l *lobby) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	l.tiles = map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.instructions = append(l.instructions, scanner.Text())
	}

	return nil
}

func (l *lobby) layout() {
	for _, instruction := range l.instructions {
		north := false
		south := false
		x, y, z := 0, 0, 0
		for i := 0; i < len(instruction); i++ {
			switch string(instruction[i]) {
			case "e":
				if north {
					x++
					z--
					north = false
				} else if south {
					z++
					y--
					south = false
				} else {
					x++
					y--
				}
			case "w":
				if north {
					y++
					z--
					north = false
				} else if south {
					z++
					x--
					south = false
				} else {
					x--
					y++
				}
			case "n":
				north = true
			case "s":
				south = true
			}
		}
		tileID := fmt.Sprintf("%d,%d,%d", x, y, z)
		l.tiles[tileID] = !l.tiles[tileID]
	}
}

func (l *lobby) countTiles() (black, white int) {
	for _, t := range l.tiles {
		if t {
			black++
		} else {
			white++
		}
	}

	return black, white
}

func (l *lobby) dailyFlip(days int) {
	for i := 0; i < days; i++ {
		fmt.Println("Day", i+1)
		newGrid := map[string]bool{}

		xMin, yMin, zMin, xMax, yMax, zMax := 0, 0, 0, 0, 0, 0
		for index := range l.tiles {
			var x, y, z int
			fmt.Sscanf(index, "%d,%d,%d", &x, &y, &z)
			if x < xMin {
				xMin = x
			}
			if x > xMax {
				xMax = x
			}

			if y < yMin {
				yMin = y
			}

			if y > yMax {
				yMax = y
			}

			if z < zMin {
				zMin = z
			}

			if z > zMax {
				zMax = z
			}
		}

		for x := xMin - 1; x <= xMax+1; x++ {
			for y := yMin - 1; y <= yMax+1; y++ {
				for z := zMin - 1; z <= zMax+1; z++ {
					blackNeighbours := 0
					if l.tiles[fmt.Sprintf("%d,%d,%d", x+1, y, z-1)] {
						blackNeighbours++
					}
					if l.tiles[fmt.Sprintf("%d,%d,%d", x+1, y-1, z)] {
						blackNeighbours++
					}
					if l.tiles[fmt.Sprintf("%d,%d,%d", x, y-1, z+1)] {
						blackNeighbours++
					}
					if l.tiles[fmt.Sprintf("%d,%d,%d", x-1, y, z+1)] {
						blackNeighbours++
					}
					if l.tiles[fmt.Sprintf("%d,%d,%d", x-1, y+1, z)] {
						blackNeighbours++
					}
					if l.tiles[fmt.Sprintf("%d,%d,%d", x, y+1, z-1)] {
						blackNeighbours++
					}

					index := fmt.Sprintf("%d,%d,%d", x, y, z)
					if l.tiles[index] {
						if blackNeighbours > 0 && blackNeighbours < 3 {
							newGrid[index] = true
						}
					} else {
						if blackNeighbours == 2 {
							newGrid[index] = true
						}
					}
				}
			}
		}
		l.tiles = newGrid
	}
}

// PartOne How many tiles are left black
func PartOne() string {
	l := lobby{}
	l.load("twentyfour/input.txt")

	l.layout()
	black, _ := l.countTiles()

	return fmt.Sprintf("There are %d black tiles", black)
}

// PartTwo How many tiles are black after 100 days
func PartTwo() string {
	l := lobby{}
	l.load("twentyfour/input.txt")

	l.layout()
	l.dailyFlip(100)
	black, _ := l.countTiles()

	return fmt.Sprintf("There are %d black tiles", black)
}
