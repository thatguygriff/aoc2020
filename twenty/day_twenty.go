package twenty

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type tile struct {
	id        int
	something [][]bool

	edges   []string
	matches map[int]bool
}

type message struct {
	tiles []tile
	grid  [][]*tile

	image [][]bool
}

func (t *tile) edge(side string) string {
	border := ""
	switch side {
	case "top":
		for _, c := range t.something[0] {
			if c {
				border += "#"
			} else {
				border += "."
			}
		}
	case "bottom":
		for _, c := range t.something[9] {
			if c {
				border += "#"
			} else {
				border += "."
			}
		}
	case "left":
		for _, row := range t.something {
			if row[0] {
				border += "#"
			} else {
				border += "."
			}
		}
	case "right":
		for _, row := range t.something {
			if row[9] {
				border += "#"
			} else {
				border += "."
			}
		}
	}

	return border
}

func (t *tile) flip() {
	flipped := [][]bool{}
	flipped = make([][]bool, 10)
	for i := 0; i < 10; i++ {
		flipped[i] = make([]bool, 10)
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			flipped[y][x] = t.something[9-y][9-x]
		}
	}

	t.something = flipped
}

func (t *tile) flipY() {
	flipped := [][]bool{}
	flipped = make([][]bool, 10)
	for i := 0; i < 10; i++ {
		flipped[i] = make([]bool, 10)
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			flipped[y][x] = t.something[9-y][x]
		}
	}

	t.something = flipped
}

func (t *tile) rotate() {
	rotated := [][]bool{}
	rotated = make([][]bool, 10)
	for i := 0; i < 10; i++ {
		rotated[i] = make([]bool, 10)
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			rotated[x][9-y] = t.something[y][x]
		}
	}

	t.something = rotated
}

func (m *message) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	m.tiles = []tile{}

	scanner := bufio.NewScanner(file)
	var t *tile
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Tile") {
			t = &tile{}
			count, err := fmt.Sscanf(scanner.Text(), "Tile %4d:", &t.id)

			if count != 1 || err != nil {
				return fmt.Errorf("Unable to parse %s", scanner.Text())
			}
			continue
		}

		if scanner.Text() == "" {
			m.tiles = append(m.tiles, *t)
			t = nil
			continue
		}

		if t != nil {
			row := []bool{}
			for _, c := range scanner.Text() {
				switch string(c) {
				case ".":
					row = append(row, false)
				case "#":
					row = append(row, true)
				}
			}

			t.something = append(t.something, row)
		}
	}

	if t != nil {
		m.tiles = append(m.tiles, *t)
	}

	for i := 0; i < len(m.tiles); i++ {
		m.tiles[i].matches = map[int]bool{}
		m.tiles[i].edges = []string{}
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("top"))
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("bottom"))
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("left"))
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("right"))
		m.tiles[i].flip()
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("top"))
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("bottom"))
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("left"))
		m.tiles[i].edges = append(m.tiles[i].edges, m.tiles[i].edge("right"))
		m.tiles[i].flip()
	}

	gridSize := int(math.Sqrt(float64(len(m.tiles))))
	m.grid = make([][]*tile, gridSize)
	for i := 0; i < gridSize; i++ {
		m.grid[i] = make([]*tile, gridSize)
	}

	m.image = make([][]bool, gridSize*8)
	for i := 0; i < gridSize*8; i++ {
		m.image[i] = make([]bool, gridSize*8)
	}

	return nil
}

func (m *message) findCorners() []tile {
	// Compute neighbours
	for i := 0; i < len(m.tiles); i++ {
		for j := 0; j < len(m.tiles); j++ {
			if i == j {
				continue
			}

			for _, iEdge := range m.tiles[i].edges {
				for _, jEdge := range m.tiles[j].edges {
					if iEdge == jEdge {
						m.tiles[i].matches[m.tiles[j].id] = true
					}
				}
			}
		}
	}

	corners := []tile{}
	// Find neighbours with only 2 matches
	for i := 0; i < len(m.tiles); i++ {
		if len(m.tiles[i].matches) == 2 {
			corners = append(corners, m.tiles[i])
		}
	}

	return corners
}

func (m *message) getTile(id int) *tile {
	for _, t := range m.tiles {
		if t.id == id {
			return &t
		}
	}
	return nil
}

func (m *message) inGrid(id int) bool {
	for y := 0; y < len(m.grid); y++ {
		for x := 0; x < len(m.grid); x++ {
			if m.grid[y][x] != nil {
				if id == m.grid[y][x].id {
					return true
				}
			}
		}
	}

	return false
}

func (m *message) layout() bool {
	corners := m.findCorners()
	for i := range corners {
		m.grid[0][0] = &corners[i]
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		m.grid[0][0].flip()
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		if m.placeNeighbours(0, 0) {
			return true
		}

		m.grid[0][0] = &corners[i]
		m.grid[0][0].rotate()
		if m.placeNeighbours(0, 0) {
			return true
		}
	}
	return false
}

func (m *message) placeNeighbours(x, y int) bool {
	allGood := true

	if x+1 >= len(m.grid) && y+1 >= len(m.grid) {
		return true
	}

	possibleTiles := []tile{}
	for id := range m.grid[y][x].matches {
		if !m.inGrid(id) {
			possibleTiles = append(possibleTiles, *m.getTile(id))
		}
	}

	if x+1 < len(m.grid) {
		if m.grid[y][x+1] == nil {
			right := m.grid[y][x].edge("right")
			found := false
			for _, p := range possibleTiles {
				for _, match := range p.edges {
					if match == right {
						found = true
						rfound := false
						for i := 0; i < 4; i++ {
							if right == p.edge("left") {
								found = true
								rfound = true
								m.grid[y][x+1] = &p
								allGood = m.placeNeighbours(x+1, y)
								break
							}

							p.rotate()
						}
						if rfound {
							break
						}

						p.flipY()
						for i := 0; i < 4; i++ {
							if right == p.edge("left") {
								found = true
								m.grid[y][x+1] = &p
								allGood = m.placeNeighbours(x+1, y)
								break
							}

							p.rotate()
						}
					}
				}
				if found {
					break
				}
			}
			allGood = found
		} else {
			allGood = m.grid[y][x].edge("right") == m.grid[y][x+1].edge("left")
		}
	}

	if y+1 < len(m.grid) {
		bottom := m.grid[y][x].edge("bottom")
		found := false
		for _, p := range possibleTiles {
			for _, match := range p.edges {
				if match == bottom {
					found = true
					rfound := false
					for i := 0; i < 4; i++ {
						if bottom == p.edge("top") {
							found = true
							rfound = true
							m.grid[y+1][x] = &p
							allGood = m.placeNeighbours(x, y+1)
							break
						}

						p.rotate()
					}
					if rfound {
						break
					}

					p.flipY()
					for i := 0; i < 4; i++ {
						if bottom == p.edge("top") {
							found = true
							m.grid[y+1][x] = &p
							allGood = m.placeNeighbours(x, y+1)
							break
						}

						p.rotate()
					}
				}
			}
			if found {
				break
			}

			allGood = false
		}
	}

	if !allGood {
		m.grid[y][x] = nil
	}

	return allGood
}

func (m *message) combineTiles() bool {
	for y := 0; y < len(m.grid); y++ {
		for x := 0; x < len(m.grid); x++ {
			if m.grid[y][x] == nil {
				return false
			}
		}
	}

	for y := 0; y < len(m.grid); y++ {
		for x := 0; x < len(m.grid); x++ {
			for tY := 1; tY < len(m.grid[y][x].something)-1; tY++ {
				for tX := 1; tX < len(m.grid[y][x].something[tY])-1; tX++ {
					imageY := (y * (len(m.grid[y][x].something) - 2)) + (tY - 1)
					imageX := (x * (len(m.grid[y][x].something[tY]) - 2)) + (tX - 1)
					m.image[imageY][imageX] = m.grid[y][x].something[tY][tX]
				}
			}
		}
	}

	return true
}

func (m *message) rotateImage() {
	rotated := [][]bool{}
	rotated = make([][]bool, len(m.image))
	for i := 0; i < len(m.image); i++ {
		rotated[i] = make([]bool, len(m.image))
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			rotated[x][9-y] = m.image[y][x]
		}
	}

	m.image = rotated
}

func (m *message) flipImage() {
	flipped := [][]bool{}
	flipped = make([][]bool, len(m.image))
	for i := 0; i < len(m.image); i++ {
		flipped[i] = make([]bool, len(m.image))
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			flipped[y][x] = m.image[9-y][x]
		}
	}

	m.image = flipped
}

func (m *message) printGrid() {
	output := ""
	for y := 0; y < len(m.grid); y++ {
		for tY := 0; tY < 10; tY++ {
			for x := 0; x < len(m.grid[y]); x++ {
				for tX := 0; tX < 10; tX++ {
					if m.grid[y][x].something[tY][tX] {
						output += "#"
					} else {
						output += "."
					}
				}
			}
			output += "\n"
		}
	}

	fmt.Printf(output)
}

func (m *message) printImage() {
	output := ""
	for _, row := range m.image {
		for _, val := range row {
			if val {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}

	fmt.Printf(output)
}

func (m *message) alignMonsters() bool {
	for i := 0; i < 4; i++ {
		count := m.seaMonsters()
		if count > 0 {
			return true
		}

		m.rotateImage()
	}
	m.flipImage()
	for i := 0; i < 4; i++ {
		count := m.seaMonsters()
		if count > 0 {
			return true
		}

		m.rotateImage()
	}
	return false
}

func (m *message) cornerProduct() int {
	corners := m.findCorners()
	if len(corners) != 4 {
		return -1
	}

	product := 1

	for _, tile := range corners {
		product *= tile.id
	}

	return product
}

func (m *message) seaMonsters() int {
	monsters := 0
	for y := 0; y < len(m.image)-2; y++ {
		for x := 0; x < len(m.image[0])-19; x++ {
			// Check if this is a valid seamonster
			if m.image[y][x+18] &&
				m.image[y+1][x] &&
				m.image[y+1][x+5] &&
				m.image[y+1][x+6] &&
				m.image[y+1][x+11] &&
				m.image[y+1][x+12] &&
				m.image[y+1][x+17] &&
				m.image[y+1][x+18] &&
				m.image[y+1][x+19] &&
				m.image[y+2][x+1] &&
				m.image[y+2][x+4] &&
				m.image[y+2][x+7] &&
				m.image[y+2][x+10] &&
				m.image[y+2][x+13] &&
				m.image[y+2][x+16] {
				monsters++
			}
		}
	}
	return monsters
}

func (m *message) countRoughness() int {
	monsters := m.seaMonsters() * 15
	total := 0
	for y := 0; y < len(m.image); y++ {
		for x := 0; x < len(m.image[0]); x++ {
			if m.image[y][x] {
				total++
			}
		}
	}

	return total - monsters
}

// PartOne Find the product of the ids of the four corners
func PartOne() string {
	m := message{}
	if err := m.load("twenty/input.txt"); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The product of the four corners is %d", m.cornerProduct())
}

// PartTwo How many # are not seamonsters.
func PartTwo() string {
	m := message{}
	if err := m.load("twenty/input.txt"); err != nil {
		return err.Error()
	}

	m.layout()
	m.combineTiles()
	m.alignMonsters()

	return fmt.Sprintf("There are %d # that are not seamonsters", m.countRoughness())
}
