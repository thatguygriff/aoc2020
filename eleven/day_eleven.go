package eleven

import (
	"bufio"
	"fmt"
	"os"
)

const (
	empty    = "L"
	occupied = "#"
	floor    = "."
)

type waitingArea struct {
	seats       [][]string
	rows, width int
}

func waitingRoom(filename string) *waitingArea {
	w := &waitingArea{}
	w.load(filename)
	return w
}

func (w *waitingArea) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if w.rows == 0 {
			w.width = len(scanner.Text())
		}

		row := make([]string, w.width)
		for position, option := range scanner.Text() {
			row[position] = string(option)
		}
		w.seats = append(w.seats, row)
		w.rows++
	}

	return nil
}

func (w *waitingArea) print() string {
	layout := ""
	for i := 0; i < w.rows; i++ {
		for j := 0; j < w.width; j++ {
			layout += w.seats[i][j]
		}
		layout += "\n"
	}
	return layout
}

func (w *waitingArea) seat(visibility bool) bool {
	changed := false

	new := w.cloneSeats()

	for i := 0; i < w.rows; i++ {
		for j := 0; j < w.width; j++ {
			switch w.seats[i][j] {
			case empty:
				if visibility {
					if w.shouldSeatVisibly(i, j) {
						new[i][j] = occupied
						changed = true
					}
				} else {
					if w.shouldSeat(i, j) {
						new[i][j] = occupied
						changed = true
					}
				}
			case occupied:
				if visibility {
					if w.shouldLeaveVisibly(i, j) {
						new[i][j] = empty
						changed = true
					}
				} else {
					if w.shouldLeave(i, j) {
						new[i][j] = empty
						changed = true
					}
				}
			case floor:
				continue
			}
		}
	}
	w.seats = new

	return changed
}

func (w *waitingArea) shouldSeat(i, j int) bool {
	okay := true
	ahead := i - 1
	behind := i + 1
	left := j - 1
	right := j + 1

	// The row ahead of the seat
	if ahead > -1 {
		if left > -1 {
			if w.seats[ahead][left] == occupied {
				okay = false
			}
		}

		if w.seats[ahead][j] == occupied {
			okay = false
		}

		if right < w.width {
			if w.seats[ahead][right] == occupied {
				okay = false
			}
		}
	}

	// The current row
	if left > -1 {
		if w.seats[i][left] == occupied {
			okay = false
		}
	}

	if right < w.width {
		if w.seats[i][right] == occupied {
			okay = false
		}
	}

	// The row behind a seat
	if behind < w.rows {
		if left > -1 {
			if w.seats[behind][left] == occupied {
				okay = false
			}
		}

		if w.seats[behind][j] == occupied {
			okay = false
		}

		if right < w.width {
			if w.seats[behind][right] == occupied {
				okay = false
			}
		}
	}

	return okay
}

func (w *waitingArea) shouldSeatVisibly(i, j int) bool {
	okay := true
	search := true
	ahead := i - 1
	behind := i + 1
	left := j - 1
	right := j + 1

	// The row ahead of the seat
	if ahead > -1 {
		leftDiagonalAhead := ahead
		leftDiagonalLeft := left
		search = true
		for search && leftDiagonalLeft > -1 && leftDiagonalAhead > -1 {
			switch w.seats[leftDiagonalAhead][leftDiagonalLeft] {
			case occupied:
				okay = false
				search = false
			case empty:
				search = false
			case floor:
				leftDiagonalAhead--
				leftDiagonalLeft--
			}
		}

		straightAhead := ahead
		search = true
		for search && straightAhead > -1 {
			switch w.seats[straightAhead][j] {
			case occupied:
				okay = false
				search = false
			case empty:
				search = false
			case floor:
				straightAhead--
			}
		}

		rightDiagonalAhead := ahead
		rightDiagonalRight := right
		search = true
		for search && rightDiagonalRight < w.width && rightDiagonalAhead > -1 {
			switch w.seats[rightDiagonalAhead][rightDiagonalRight] {
			case occupied:
				okay = false
				search = false
			case empty:
				search = false
			case floor:
				rightDiagonalAhead--
				rightDiagonalRight++
			}
		}
	}

	// The current row
	currentLeft := left
	search = true
	for search && currentLeft > -1 {
		switch w.seats[i][currentLeft] {
		case occupied:
			okay = false
			search = false
		case empty:
			search = false
		case floor:
			currentLeft--
		}
	}

	currentRight := right
	search = true
	for search && currentRight < w.width {
		switch w.seats[i][currentRight] {
		case occupied:
			okay = false
			search = false
		case empty:
			search = false
		case floor:
			currentRight++
		}
	}

	// The row behind a seat
	if behind < w.rows {
		leftDiagonalBehind := behind
		leftDiagonalLeft := left
		search = true
		for search && leftDiagonalLeft > -1 && leftDiagonalBehind < w.rows {
			switch w.seats[leftDiagonalBehind][leftDiagonalLeft] {
			case occupied:
				okay = false
				search = false
			case empty:
				search = false
			case floor:
				leftDiagonalBehind++
				leftDiagonalLeft--
			}
		}

		straightBehind := behind
		search = true
		for search && straightBehind < w.rows {
			switch w.seats[straightBehind][j] {
			case occupied:
				okay = false
				search = false
			case empty:
				search = false
			case floor:
				straightBehind++
			}
		}

		rightDiagonalBehind := behind
		rightDiagonalRight := right
		search = true
		for search && rightDiagonalRight < w.width && rightDiagonalBehind < w.rows {
			switch w.seats[rightDiagonalBehind][rightDiagonalRight] {
			case occupied:
				okay = false
				search = false
			case empty:
				search = false
			case floor:
				rightDiagonalBehind++
				rightDiagonalRight++
			}
		}
	}

	return okay
}

func (w *waitingArea) shouldLeave(i, j int) bool {
	ahead := i - 1
	behind := i + 1
	left := j - 1
	right := j + 1

	occupiedSeats := 0

	// The row ahead of the seat
	if ahead > -1 {
		if left > -1 {
			if w.seats[ahead][left] == occupied {
				occupiedSeats++
			}
		}

		if w.seats[ahead][j] == occupied {
			occupiedSeats++
		}

		if right < w.width {
			if w.seats[ahead][right] == occupied {
				occupiedSeats++
			}
		}
	}

	// The current row
	if left > -1 {
		if w.seats[i][left] == occupied {
			occupiedSeats++
		}
	}

	if right < w.width {
		if w.seats[i][right] == occupied {
			occupiedSeats++
		}
	}

	// The row behind a seat
	if behind < w.rows {
		if left > -1 {
			if w.seats[behind][left] == occupied {
				occupiedSeats++
			}
		}

		if w.seats[behind][j] == occupied {
			occupiedSeats++
		}

		if right < w.width {
			if w.seats[behind][right] == occupied {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats >= 4
}

func (w *waitingArea) shouldLeaveVisibly(i, j int) bool {
	ahead := i - 1
	behind := i + 1
	left := j - 1
	right := j + 1
	search := true
	occupiedSeats := 0

	// The row ahead of the seat
	if ahead > -1 {
		leftDiagonalAhead := ahead
		leftDiagonalLeft := left
		search = true
		for search && leftDiagonalLeft > -1 && leftDiagonalAhead > -1 {
			switch w.seats[leftDiagonalAhead][leftDiagonalLeft] {
			case occupied:
				occupiedSeats++
				search = false
			case empty:
				search = false
			case floor:
				leftDiagonalAhead--
				leftDiagonalLeft--
			}
		}

		straightAhead := ahead
		search = true
		for search && straightAhead > -1 {
			switch w.seats[straightAhead][j] {
			case occupied:
				occupiedSeats++
				search = false
			case empty:
				search = false
			case floor:
				straightAhead--
			}
		}

		rightDiagonalAhead := ahead
		rightDiagonalRight := right
		search = true
		for search && rightDiagonalRight < w.width && rightDiagonalAhead > -1 {
			switch w.seats[rightDiagonalAhead][rightDiagonalRight] {
			case occupied:
				occupiedSeats++
				search = false
			case empty:
				search = false
			case floor:
				rightDiagonalAhead--
				rightDiagonalRight++
			}
		}
	}

	// The current row
	currentLeft := left
	search = true
	for search && currentLeft > -1 {
		switch w.seats[i][currentLeft] {
		case occupied:
			occupiedSeats++
			search = false
		case empty:
			search = false
		case floor:
			currentLeft--
		}
	}

	currentRight := right
	search = true
	for search && currentRight < w.width {
		switch w.seats[i][currentRight] {
		case occupied:
			occupiedSeats++
			search = false
		case empty:
			search = false
		case floor:
			currentRight++
		}
	}

	// The row behind a seat
	if behind < w.rows {
		leftDiagonalBehind := behind
		leftDiagonalLeft := left
		search = true
		for search && leftDiagonalLeft > -1 && leftDiagonalBehind < w.rows {
			switch w.seats[leftDiagonalBehind][leftDiagonalLeft] {
			case occupied:
				occupiedSeats++
				search = false
			case empty:
				search = false
			case floor:
				leftDiagonalBehind++
				leftDiagonalLeft--
			}
		}

		straightBehind := behind
		search = true
		for search && straightBehind < w.rows {
			switch w.seats[straightBehind][j] {
			case occupied:
				occupiedSeats++
				search = false
			case empty:
				search = false
			case floor:
				straightBehind++
			}
		}

		rightDiagonalBehind := behind
		rightDiagonalRight := right
		search = true
		for search && rightDiagonalRight < w.width && rightDiagonalBehind < w.rows {
			switch w.seats[rightDiagonalBehind][rightDiagonalRight] {
			case occupied:
				occupiedSeats++
				search = false
			case empty:
				search = false
			case floor:
				rightDiagonalBehind++
				rightDiagonalRight++
			}
		}
	}

	return occupiedSeats >= 5
}

func (w *waitingArea) simulate(visible bool) int {
	// fmt.Println(w.print())
	for w.seat(visible) {
		// fmt.Println(w.print())
	}

	count := 0
	for i := 0; i < w.rows; i++ {
		for j := 0; j < w.width; j++ {
			if w.seats[i][j] == occupied {
				count++
			}
		}
	}

	return count
}

func (w *waitingArea) cloneSeats() [][]string {
	new := make([][]string, w.rows)

	for i, row := range w.seats {
		new[i] = make([]string, w.width)
		for j, seat := range row {
			new[i][j] = seat
		}
	}

	return new
}

// PartOne How many seats are occupied in a stable seating arrangement
func PartOne() string {
	w := waitingRoom("eleven/input.txt")

	return fmt.Sprintf("A stable seating arrangement has %d seats", w.simulate(false))
}

// PartTwo How many seats are occupied in a stable seating arrangement based on visibility
func PartTwo() string {
	w := waitingRoom("eleven/input.txt")

	return fmt.Sprintf("A stable seating arrangement has %d seats", w.simulate(true))
}
