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

// PartOne How many tiles are left black
func PartOne() string {
	l := lobby{}
	l.load("twentyfour/input.txt")

	l.layout()
	black, _ := l.countTiles()

	return fmt.Sprintf("There are %d black tiles", black)
}
