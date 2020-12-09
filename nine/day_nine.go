package nine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type xmas struct {
	preamble int
	values   []int
}

func newXmas(filename string, preamble int) (*xmas, error) {
	x := &xmas{}
	if err := x.load(filename, preamble); err != nil {
		return nil, err
	}

	return x, nil
}

func (x *xmas) load(filename string, preamble int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		x.values = append(x.values, v)
	}

	x.preamble = preamble

	return nil
}

func (x *xmas) detectVulnerability() (int, error) {
	index := x.preamble

	if index > len(x.values) {
		return 0, fmt.Errorf("Preamble is larger than the range of inputs")
	}

	var match bool

	for index < len(x.values) {
		searchable := x.values[index-x.preamble : index]
		match = false

		for i := 0; i < len(searchable); i++ {
			for j := i + 1; j < len(searchable); j++ {
				if searchable[i] == searchable[j] {
					continue
				}

				if searchable[i]+searchable[j] == x.values[index] {
					match = true
				}
			}
		}

		if !match {
			return x.values[index], nil
		}
		index++
	}

	return 0, fmt.Errorf("Did not find a vulnerable value")
}

func (x *xmas) computeWeakness() (int, error) {
	v, err := x.detectVulnerability()
	if err != nil {
		return 0, nil
	}

	for i := 0; i < len(x.values); i++ {
		sum := x.values[i]
		min := x.values[i]
		max := x.values[i]

		if x.values[i] == v {
			return 0, fmt.Errorf("Did not find encryption weakness before %d", v)
		}

		for j := i + 1; j < len(x.values); j++ {
			sum += x.values[j]
			if x.values[j] < min {
				min = x.values[j]
			}

			if x.values[j] > max {
				max = x.values[j]
			}

			// We found the contiguous range of values!
			if sum == v {
				return min + max, nil
			}

			// This means that the contiguous run starting at i was a bust
			if sum > v {
				break
			}
		}
	}

	return 0, nil
}

// PartOne Find the first vulnerable number in the xmas chain
func PartOne() string {
	x, _ := newXmas("nine/input.txt", 25)

	v, _ := x.detectVulnerability()

	return fmt.Sprintf("The first vulnerable value in the XMAS input is %d", v)
}

// PartTwo Find the encryption weaknes in the xmas chain
func PartTwo() string {
	x, _ := newXmas("nine/input.txt", 25)

	w, _ := x.computeWeakness()

	return fmt.Sprintf("The encryption weakness in the XMAS input is %d", w)
}
