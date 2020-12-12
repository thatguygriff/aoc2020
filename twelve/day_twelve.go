package twelve

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type instruction struct {
	action string
	value  int
}

type waypoint struct {
	x, y int
}

type boat struct {
	orientation int
	orders      []instruction
	x, y        int
	waypoint    waypoint
}

func (b *boat) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		amount, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			return err
		}

		b.orders = append(b.orders, instruction{
			action: scanner.Text()[:1],
			value:  amount,
		})
	}

	return nil
}

func loadBoat(filename string) (*boat, error) {
	b := &boat{
		waypoint: waypoint{
			x: 10,
			y: 1,
		},
	}

	err := b.load(filename)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *boat) advance(distance int) {
	if b.orientation > 315 || b.orientation < 45 {
		b.x += distance
		return
	}

	if b.orientation > 45 && b.orientation < 135 {
		b.y += distance
		return
	}

	if b.orientation > 135 && b.orientation < 225 {
		b.x -= distance
	}

	if b.orientation > 225 && b.orientation < 315 {
		b.y -= distance
	}
}

func (b *boat) navigate() int {
	for _, o := range b.orders {
		switch o.action {
		case "F":
			b.advance(o.value)
		case "N":
			b.y += o.value
		case "S":
			b.y -= o.value
		case "E":
			b.x += o.value
		case "W":
			b.x -= o.value
		case "L":
			b.orientation += o.value
		case "R":
			b.orientation -= o.value
		}

		for b.orientation >= 360 {
			b.orientation -= 360
		}

		for b.orientation < 0 {
			b.orientation += 360
		}
	}

	return int(math.Abs(float64(b.x))) + int(math.Abs(float64(b.y)))
}

func (b *boat) rotate(degrees int) {
	newX, newY := b.waypoint.x, b.waypoint.y

	if degrees == 90 || degrees == -270 {
		newX = b.waypoint.y * -1
		newY = b.waypoint.x
	}

	if degrees == 180 || degrees == -180 {
		newX = b.waypoint.x * -1
		newY = b.waypoint.y * -1
	}

	if degrees == 270 || degrees == -90 {
		newX = b.waypoint.y
		newY = b.waypoint.x * -1
	}

	b.waypoint.x = newX
	b.waypoint.y = newY
}

func (b *boat) waypointNavigate() int {
	for _, o := range b.orders {
		switch o.action {
		case "F":
			b.x += o.value * b.waypoint.x
			b.y += o.value * b.waypoint.y
		case "N":
			b.waypoint.y += o.value
		case "S":
			b.waypoint.y -= o.value
		case "E":
			b.waypoint.x += o.value
		case "W":
			b.waypoint.x -= o.value
		case "L":
			b.rotate(o.value)
		case "R":
			b.rotate(o.value * -1)
		}

		for b.orientation >= 360 {
			b.orientation -= 360
		}

		for b.orientation < 0 {
			b.orientation += 360
		}
	}
	return int(math.Abs(float64(b.x))) + int(math.Abs(float64(b.y)))
}

// PartOne What is the Manhatten distance that the boat travelled?
func PartOne() string {
	b, err := loadBoat("twelve/input.txt")
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The boat travelled a manhattan distance of %d", b.navigate())
}

// PartTwo What is the Manhatten distance that the boat travelled using a waypoint?
func PartTwo() string {
	b, err := loadBoat("twelve/input.txt")
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The boat travelled a manhattan distance of %d using a waypoint", b.waypointNavigate())
}
