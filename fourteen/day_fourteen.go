package fourteen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	mem     int
	pointer int

	mask   string
	isMask bool
}

type computer struct {
	memory map[int]uint64
	m      map[int]bool

	mask  string
	ones  uint64
	zeros uint64

	input    []command
	executed bool
}

func (c *computer) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " = ")
		if len(inputs) != 2 {
			return fmt.Errorf("Unable to parse input %q", scanner.Text())
		}

		if inputs[0] == "mask" {
			c.input = append(c.input, command{
				isMask: true,
				mask:   inputs[1],
			})
		} else {
			var pointer, mem int
			count, err := fmt.Sscanf(inputs[0], "mem[%d]", &pointer)
			if count != 1 || err != nil {
				return fmt.Errorf("Unable to parse memory location from %s", inputs[0])
			}
			mem, err = strconv.Atoi(inputs[1])

			c.input = append(c.input, command{
				isMask:  false,
				mem:     mem,
				pointer: pointer,
			})
		}
	}

	c.memory = make(map[int]uint64, len(c.input))
	return nil
}

func newComputer(filename string) *computer {
	c := &computer{}
	c.mask = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	c.executed = false
	if err := c.load(filename); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return c
}

func (c *computer) execute() {
	c.executed = true

	for _, input := range c.input {
		if input.isMask {
			c.mask = input.mask
			c.ones = 0
			c.zeros = 0

			for i := 0; i < len(input.mask); i++ {
				switch string(input.mask[len(input.mask)-1-i]) {
				case "0":
					c.zeros += 1 << i
				case "1":
					c.ones += 1 << i
				}
			}

			continue
		}

		value := uint64(input.mem)
		value |= c.ones
		value &^= c.zeros
		c.memory[input.pointer] = value
	}
}

func (c *computer) executeV2() {
	c.executed = true

	for _, input := range c.input {
		if input.isMask {
			c.mask = input.mask
			c.ones = 0
			c.zeros = 0
			for i := 0; i < len(input.mask); i++ {
				switch string(input.mask[len(input.mask)-1-i]) {
				case "0":
					c.zeros += 1 << i
				case "1":
					c.ones += 1 << i
				}
			}

			continue
		}

		value := uint64(input.mem)
		location := uint64(input.pointer) | c.ones

		for _, p := range c.memMask(location) {
			c.memory[p] = value
		}
	}
}

func (c *computer) memMask(p uint64) []int {
	locations := []int{0}

	for i := 0; i < len(c.mask); i++ {
		switch string(c.mask[len(c.mask)-1-i]) {
		case "X":
			for j, l := range locations {
				locations = append(locations, int(l|(1<<i)))
				locations[j] = int(l &^ (1 << i))
			}
		case "1":
			for j := 0; j < len(locations); j++ {
				locations[j] = int(locations[j] | (1 << i))
			}
		case "0":
			for j := 0; j < len(locations); j++ {
				locations[j] = locations[j] | int(p&(1<<i))
			}
		}

	}
	return locations
}

func (c *computer) sum(version int) int {
	if !c.executed {
		switch version {
		case 1:
			c.execute()
		case 2:
			c.executeV2()
		}
	}

	sum := 0
	for _, v := range c.memory {
		sum += int(v)
	}

	return sum
}

// PartOne Execute the initialization program
func PartOne() string {
	c := newComputer("fourteen/input.txt")
	if c == nil {
		return "Unable to load program"
	}

	return fmt.Sprintf("The memory value after execution is %d", c.sum(1))
}

// PartTwo Execute the initialization program with version 2
func PartTwo() string {
	c := newComputer("fourteen/input.txt")
	if c == nil {
		return "Unable to load program"
	}

	return fmt.Sprintf("The memory value after execution is %d", c.sum(2))
}
